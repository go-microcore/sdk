package adapter

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"go.microcore.dev/framework/transport/http"
	"go.microcore.dev/framework/transport/http/client"
)

type Config struct {
	HttpClientManager    client.Manager
	UsersServiceEndpoint string
}

func New(config *Config) Interface {
	return &adapter{
		config.HttpClientManager,
		config.UsersServiceEndpoint,
	}
}

type adapter struct {
	httpClientManager    client.Manager
	usersServiceEndpoint string
}

func (a *adapter) Signin(ctx context.Context, data SigninData) (*SigninResult, error) {
	// Encode body json
	body, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error parsing request body: %v", err)
	}

	// Send service request
	res, err := a.httpClientManager.Request(
		a.usersServiceEndpoint+"/users/signin",
		client.WithRequestMethod(http.MethodPost),
		client.WithRequestBody(body),
		client.WithRequestContext(ctx),
	)
	if err != nil {
		return nil, fmt.Errorf("service %s unavailable: %v", a.usersServiceEndpoint, err)
	}

	// Check success status code
	if res.StatusCode() == 200 {
		var response signinResponse
		if err := json.Unmarshal(res.Body(), &response); err != nil {
			return nil, fmt.Errorf("error parsing response body: %v", err)
		}
		result := SigninResult(response)
		return &result, nil
	}

	// Response message
	errMessage := string(res.Body())

	// Errors map
	var errMap = map[string]error{
		"bad_request:invalid_login":        ErrInvalidLogin,
		"bad_request:invalid_password":     ErrInvalidPassword,
		"unauthorized:invalid_credentials": ErrInvalidCredentials,
	}

	// Parse errors
	if err, ok := errMap[errMessage]; ok {
		return nil, err
	}

	return nil, fmt.Errorf("unexpected response: status code: %d, message: %s", res.StatusCode(), errMessage)
}

func (a *adapter) Signup(ctx context.Context, data SignupData) (*SignupResult, error) {
	// Encode body json
	body, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error parsing request body: %v", err)
	}

	// Send service request
	res, err := a.httpClientManager.Request(
		a.usersServiceEndpoint+"/users/signup",
		client.WithRequestMethod(http.MethodPost),
		client.WithRequestBody(body),
		client.WithRequestContext(ctx),
	)
	if err != nil {
		return nil, fmt.Errorf("service %s unavailable: %v", a.usersServiceEndpoint, err)
	}

	// Check success status code
	if res.StatusCode() == 201 {
		var response signupResponse
		if err := json.Unmarshal(res.Body(), &response); err != nil {
			return nil, fmt.Errorf("error parsing response body: %v", err)
		}
		result := SignupResult{
			Id:         response.Id,
			Created:    response.Created,
			Username:   response.Username,
			Email:      response.Email,
			Name:       response.Name,
			Role:       SignupRoleResult(response.Role),
			Mfa:        response.Mfa,
			SystemFlag: response.SystemFlag,
		}
		return &result, nil
	}

	// Response message
	errMessage := string(res.Body())

	// Errors map
	var errMap = map[string]error{
		"bad_request:invalid_name":        ErrInvalidName,
		"bad_request:invalid_username":    ErrInvalidUsername,
		"bad_request:invalid_email":       ErrInvalidEmail,
		"bad_request:invalid_password":    ErrInvalidPassword,
		"bad_request:user_exist_email":    ErrExistEmail,
		"bad_request:user_exist_username": ErrExistUsername,
	}

	// Parse errors
	if err, ok := errMap[errMessage]; ok {
		return nil, err
	}

	return nil, fmt.Errorf("unexpected response: status code: %d, message: %s", res.StatusCode(), errMessage)
}

func (a *adapter) Profile(ctx context.Context, authToken string) (*ProfileResult, error) {
	// Send service request
	res, err := a.httpClientManager.Request(
		a.usersServiceEndpoint+"/users/profile",
		client.WithRequestMethod(http.MethodGet),
		client.WithRequestContext(ctx),
		client.WithRequestHeaders(
			client.NewRequestHeader("Authorization", "Bearer "+authToken),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("service %s unavailable: %v", a.usersServiceEndpoint, err)
	}

	// Check success status code
	if res.StatusCode() == 200 {
		var response profileResponse
		if err := json.Unmarshal(res.Body(), &response); err != nil {
			return nil, fmt.Errorf("error parsing response body: %v", err)
		}
		result := ProfileResult{
			Id:         response.Id,
			Created:    response.Created,
			Username:   response.Username,
			Email:      response.Email,
			Name:       response.Name,
			Role:       ProfileRoleResult(response.Role),
			Mfa:        response.Mfa,
			SystemFlag: response.SystemFlag,
			Device:     response.Device,
		}
		return &result, nil
	}

	// Response message
	errMessage := string(res.Body())

	return nil, fmt.Errorf("unexpected response: status code: %d, message: %s", res.StatusCode(), errMessage)
}

func (a *adapter) TwoFAValidate(ctx context.Context, authToken string, data TwoFAValidateData) (*TwoFAValidateResult, error) {
	// Encode body json
	body, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error parsing request body: %v", err)
	}

	// Send service request
	res, err := a.httpClientManager.Request(
		a.usersServiceEndpoint+"/users/2fa/validate",
		client.WithRequestMethod(http.MethodPost),
		client.WithRequestBody(body),
		client.WithRequestContext(ctx),
		client.WithRequestHeaders(
			client.NewRequestHeader("Authorization", "Bearer "+authToken),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("service %s unavailable: %v", a.usersServiceEndpoint, err)
	}

	// Check success status code
	if res.StatusCode() == 200 {
		var response twoFAValidateResponse
		if err := json.Unmarshal(res.Body(), &response); err != nil {
			return nil, fmt.Errorf("error parsing response body: %v", err)
		}
		result := TwoFAValidateResult(response)
		return &result, nil
	}

	// Response message
	errMessage := string(res.Body())

	// Errors map
	var errMap = map[string]error{
		"bad_request:invalid_token": ErrInvalidToken,
		"bad_request:mfa_disabled":  ErrMfaDisabled,
	}

	// Parse errors
	if err, ok := errMap[errMessage]; ok {
		return nil, err
	}

	return nil, fmt.Errorf("unexpected response: status code: %d, message: %s", res.StatusCode(), errMessage)
}

func (a *adapter) TwoFASettings(ctx context.Context, authToken string, data TwoFASettingsData) (*TwoFASettingsResult, error) {
	// Encode body json
	body, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error parsing request body: %v", err)
	}

	// Send service request
	res, err := a.httpClientManager.Request(
		a.usersServiceEndpoint+"/users/2fa/settings/",
		client.WithRequestMethod(http.MethodPost),
		client.WithRequestBody(body),
		client.WithRequestContext(ctx),
		client.WithRequestHeaders(
			client.NewRequestHeader("Authorization", "Bearer "+authToken),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("service %s unavailable: %v", a.usersServiceEndpoint, err)
	}

	// Check success status code
	if res.StatusCode() == 200 {
		var response twoFASettingsResponse
		if err := json.Unmarshal(res.Body(), &response); err != nil {
			return nil, fmt.Errorf("error parsing response body: %v", err)
		}
		result := TwoFASettingsResult(response)
		return &result, nil
	}

	// Response message
	errMessage := string(res.Body())

	// Errors map
	var errMap = map[string]error{
		"bad_request:invalid_password":     ErrInvalidPassword,
		"unauthorized:invalid_credentials": ErrInvalidCredentials,
	}

	// Parse errors
	if err, ok := errMap[errMessage]; ok {
		return nil, err
	}

	return nil, fmt.Errorf("unexpected response: status code: %d, message: %s", res.StatusCode(), errMessage)
}

func (a *adapter) TwoFAEnable(ctx context.Context, authToken string, data TwoFAEnableData) error {
	// Encode body json
	body, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("error parsing request body: %v", err)
	}

	// Send service request
	res, err := a.httpClientManager.Request(
		a.usersServiceEndpoint+"/users/2fa/settings/enable",
		client.WithRequestMethod(http.MethodPost),
		client.WithRequestBody(body),
		client.WithRequestContext(ctx),
		client.WithRequestHeaders(
			client.NewRequestHeader("Authorization", "Bearer "+authToken),
		),
	)
	if err != nil {
		return fmt.Errorf("service %s unavailable: %v", a.usersServiceEndpoint, err)
	}

	// Check success status code
	if res.StatusCode() == 204 {
		return nil
	}

	// Response message
	errMessage := string(res.Body())

	// Errors map
	var errMap = map[string]error{
		"bad_request:invalid_token": ErrInvalidToken,
		"bad_request:mfa_enabled":   ErrMfaEnabled,
	}

	// Parse errors
	if err, ok := errMap[errMessage]; ok {
		return err
	}

	return fmt.Errorf("unexpected response: status code: %d, message: %s", res.StatusCode(), errMessage)
}

func (a *adapter) TwoFADisable(ctx context.Context, authToken string, data TwoFADisableData) error {
	// Encode body json
	body, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("error parsing request body: %v", err)
	}

	// Send service request
	res, err := a.httpClientManager.Request(
		a.usersServiceEndpoint+"/users/2fa/settings/disable",
		client.WithRequestMethod(http.MethodPost),
		client.WithRequestBody(body),
		client.WithRequestContext(ctx),
		client.WithRequestHeaders(
			client.NewRequestHeader("Authorization", "Bearer "+authToken),
		),
	)
	if err != nil {
		return fmt.Errorf("service %s unavailable: %v", a.usersServiceEndpoint, err)
	}

	// Check success status code
	if res.StatusCode() == 204 {
		return nil
	}

	// Response message
	errMessage := string(res.Body())

	// Errors map
	var errMap = map[string]error{
		"bad_request:invalid_password":     ErrInvalidPassword,
		"bad_request:invalid_token":        ErrInvalidToken,
		"bad_request:mfa_disabled":         ErrMfaDisabled,
		"unauthorized:invalid_credentials": ErrInvalidCredentials,
	}

	// Parse errors
	if err, ok := errMap[errMessage]; ok {
		return err
	}

	return fmt.Errorf("unexpected response: status code: %d, message: %s", res.StatusCode(), errMessage)
}

func (a *adapter) CreateUser(ctx context.Context, authToken string, data CreateUserData) (*CreateUserResult, error) {
	// Encode body json
	body, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error parsing request body: %v", err)
	}

	// Send service request
	res, err := a.httpClientManager.Request(
		a.usersServiceEndpoint+"/users/admin/",
		client.WithRequestMethod(http.MethodPost),
		client.WithRequestBody(body),
		client.WithRequestContext(ctx),
		client.WithRequestHeaders(
			client.NewRequestHeader("Authorization", "Bearer "+authToken),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("service %s unavailable: %v", a.usersServiceEndpoint, err)
	}

	// Check success status code
	if res.StatusCode() == 201 {
		var response createUserResponse
		if err := json.Unmarshal(res.Body(), &response); err != nil {
			return nil, fmt.Errorf("error parsing response body: %v", err)
		}
		result := CreateUserResult{
			Id:         response.Id,
			Created:    response.Created,
			Username:   response.Username,
			Email:      response.Email,
			Name:       response.Name,
			Role:       CreateUserRoleResult(response.Role),
			Mfa:        response.Mfa,
			SystemFlag: response.SystemFlag,
		}
		return &result, nil
	}

	// Response message
	errMessage := string(res.Body())

	// Errors map
	var errMap = map[string]error{
		"bad_request:invalid_name":        ErrInvalidName,
		"bad_request:invalid_username":    ErrInvalidUsername,
		"bad_request:invalid_password":    ErrInvalidPassword,
		"bad_request:invalid_email":       ErrInvalidEmail,
		"bad_request:invalid_roles":       ErrInvalidRoles,
		"bad_request:user_exist_username": ErrExistUsername,
		"bad_request:user_exist_email":    ErrExistEmail,
	}

	// Parse errors
	if err, ok := errMap[errMessage]; ok {
		return nil, err
	}

	return nil, fmt.Errorf("unexpected response: status code: %d, message: %s", res.StatusCode(), errMessage)
}

func (a *adapter) FilterUsers(ctx context.Context, authToken string, data FilterUsersData) ([]FilterUsersResult, error) {
	// Encode body json
	body, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error parsing request body: %v", err)
	}

	// Send service request
	res, err := a.httpClientManager.Request(
		a.usersServiceEndpoint+"/users/admin/filter",
		client.WithRequestMethod(http.MethodPost),
		client.WithRequestBody(body),
		client.WithRequestContext(ctx),
		client.WithRequestHeaders(
			client.NewRequestHeader("Authorization", "Bearer "+authToken),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("service %s unavailable: %v", a.usersServiceEndpoint, err)
	}

	// Check success status code
	if res.StatusCode() == 200 {
		var response []filterUsersResponse
		if err := json.Unmarshal(res.Body(), &response); err != nil {
			return nil, fmt.Errorf("error parsing response body: %v", err)
		}
		result := make([]FilterUsersResult, len(response))
		for index, item := range response {
			result[index] = FilterUsersResult{
				Id:         item.Id,
				Created:    item.Created,
				Username:   item.Username,
				Email:      item.Email,
				Name:       item.Name,
				Role:       FilterUsersRoleResult(item.Role),
				Mfa:        item.Mfa,
				SystemFlag: item.SystemFlag,
			}
		}
		return result, nil
	}

	// Response message
	errMessage := string(res.Body())

	return nil, fmt.Errorf("unexpected response: status code: %d, message: %s", res.StatusCode(), errMessage)
}

func (a *adapter) UpdateUser(ctx context.Context, authToken string, id string, data UpdateUserData) error {
	// Encode body json
	body, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("error parsing request body: %v", err)
	}

	// Build path
	var path strings.Builder
	path.WriteString("/users/admin/")
	path.WriteString(id)

	// Send service request
	res, err := a.httpClientManager.Request(
		a.usersServiceEndpoint+path.String(),
		client.WithRequestMethod(http.MethodPatch),
		client.WithRequestBody(body),
		client.WithRequestContext(ctx),
		client.WithRequestHeaders(
			client.NewRequestHeader("Authorization", "Bearer "+authToken),
		),
	)
	if err != nil {
		return fmt.Errorf("service %s unavailable: %v", a.usersServiceEndpoint, err)
	}

	// Check success status code
	if res.StatusCode() == 204 {
		return nil
	}

	// Response message
	errMessage := string(res.Body())

	// Errors map
	var errMap = map[string]error{
		"bad_request:invalid_name":        ErrInvalidName,
		"bad_request:invalid_username":    ErrInvalidUsername,
		"bad_request:invalid_email":       ErrInvalidEmail,
		"bad_request:invalid_roles":       ErrInvalidRoles,
		"bad_request:user_exist_username": ErrExistUsername,
		"bad_request:user_exist_email":    ErrExistEmail,
		"bad_request:user_not_found":      ErrNotFound,
	}

	// Parse errors
	if err, ok := errMap[errMessage]; ok {
		return err
	}

	return fmt.Errorf("unexpected response: status code: %d, message: %s", res.StatusCode(), errMessage)
}

func (a *adapter) DeleteUser(ctx context.Context, authToken string, id uint) error {
	// Build path
	var path strings.Builder
	path.WriteString("/users/admin/")
	path.WriteString(strconv.FormatUint(uint64(id), 10))

	// Send service request
	res, err := a.httpClientManager.Request(
		a.usersServiceEndpoint+path.String(),
		client.WithRequestMethod(http.MethodDelete),
		client.WithRequestContext(ctx),
		client.WithRequestHeaders(
			client.NewRequestHeader("Authorization", "Bearer "+authToken),
		),
	)
	if err != nil {
		return fmt.Errorf("service %s unavailable: %v", a.usersServiceEndpoint, err)
	}

	// Check success status code
	if res.StatusCode() == 204 {
		return nil
	}

	// Response message
	errMessage := string(res.Body())

	// Errors map
	var errMap = map[string]error{
		"bad_request:user_not_found": ErrNotFound,
		"bad_request:user_is_used":   ErrUserIsUsed,
	}

	// Parse errors
	if err, ok := errMap[errMessage]; ok {
		return err
	}

	return fmt.Errorf("unexpected response: status code: %d, message: %s", res.StatusCode(), errMessage)
}

func (a *adapter) CreateRole(ctx context.Context, authToken string, data CreateRoleData) (*CreateRoleResult, error) {
	// Encode body json
	body, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error parsing request body: %v", err)
	}

	// Send service request
	res, err := a.httpClientManager.Request(
		a.usersServiceEndpoint+"/users/admin/roles/",
		client.WithRequestMethod(http.MethodPost),
		client.WithRequestBody(body),
		client.WithRequestContext(ctx),
		client.WithRequestHeaders(
			client.NewRequestHeader("Authorization", "Bearer "+authToken),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("service %s unavailable: %v", a.usersServiceEndpoint, err)
	}

	// Check success status code
	if res.StatusCode() == 201 {
		var response createRoleResponse
		if err := json.Unmarshal(res.Body(), &response); err != nil {
			return nil, fmt.Errorf("error parsing response body: %v", err)
		}
		result := CreateRoleResult(response)
		return &result, nil
	}

	// Response message
	errMessage := string(res.Body())

	// Errors map
	var errMap = map[string]error{
		"bad_request:invalid_role_id":   ErrInvalidRoleId,
		"bad_request:invalid_role_name": ErrInvalidRoleName,
		"bad_request:role_exist_id":     ErrExistRoleId,
		"bad_request:role_exist_name":   ErrExistRoleName,
	}

	// Parse errors
	if err, ok := errMap[errMessage]; ok {
		return nil, err
	}

	return nil, fmt.Errorf("unexpected response: status code: %d, message: %s", res.StatusCode(), errMessage)
}

func (a *adapter) FilterRoles(ctx context.Context, authToken string, data FilterRolesData) ([]FilterRolesResult, error) {
	// Encode body json
	body, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error parsing request body: %v", err)
	}

	// Send service request
	res, err := a.httpClientManager.Request(
		a.usersServiceEndpoint+"/users/admin/roles/filter",
		client.WithRequestMethod(http.MethodPost),
		client.WithRequestBody(body),
		client.WithRequestContext(ctx),
		client.WithRequestHeaders(
			client.NewRequestHeader("Authorization", "Bearer "+authToken),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("service %s unavailable: %v", a.usersServiceEndpoint, err)
	}

	// Check success status code
	if res.StatusCode() == 200 {
		var response []filterRolesResponse
		if err := json.Unmarshal(res.Body(), &response); err != nil {
			return nil, fmt.Errorf("error parsing response body: %v", err)
		}
		result := make([]FilterRolesResult, len(response))
		for index, item := range response {
			result[index] = FilterRolesResult(item)
		}
		return result, nil
	}

	// Response message
	errMessage := string(res.Body())

	return nil, fmt.Errorf("unexpected response: status code: %d, message: %s", res.StatusCode(), errMessage)
}

func (a *adapter) UpdateRole(ctx context.Context, authToken string, id string, data UpdateRoleData) error {
	// Encode body json
	body, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("error parsing request body: %v", err)
	}

	// Build path
	var path strings.Builder
	path.WriteString("/users/admin/roles/")
	path.WriteString(id)

	// Send service request
	res, err := a.httpClientManager.Request(
		a.usersServiceEndpoint+path.String(),
		client.WithRequestMethod(http.MethodPatch),
		client.WithRequestBody(body),
		client.WithRequestContext(ctx),
		client.WithRequestHeaders(
			client.NewRequestHeader("Authorization", "Bearer "+authToken),
		),
	)
	if err != nil {
		return fmt.Errorf("service %s unavailable: %v", a.usersServiceEndpoint, err)
	}

	// Check success status code
	if res.StatusCode() == 204 {
		return nil
	}

	// Response message
	errMessage := string(res.Body())

	// Errors map
	var errMap = map[string]error{
		"bad_request:role_not_found":    ErrRoleNotFound,
		"bad_request:invalid_role_name": ErrInvalidRoleName,
		"bad_request:role_exist_name":   ErrExistRoleName,
	}

	// Parse errors
	if err, ok := errMap[errMessage]; ok {
		return err
	}

	return fmt.Errorf("unexpected response: status code: %d, message: %s", res.StatusCode(), errMessage)
}

func (a *adapter) DeleteRole(ctx context.Context, authToken string, id string) error {
	// Build path
	var path strings.Builder
	path.WriteString("/users/admin/roles/")
	path.WriteString(id)

	// Send service request
	res, err := a.httpClientManager.Request(
		a.usersServiceEndpoint+path.String(),
		client.WithRequestMethod(http.MethodDelete),
		client.WithRequestContext(ctx),
		client.WithRequestHeaders(
			client.NewRequestHeader("Authorization", "Bearer "+authToken),
		),
	)
	if err != nil {
		return fmt.Errorf("service %s unavailable: %v", a.usersServiceEndpoint, err)
	}

	// Check success status code
	if res.StatusCode() == 204 {
		return nil
	}

	// Response message
	errMessage := string(res.Body())

	// Errors map
	var errMap = map[string]error{
		"bad_request:role_not_found": ErrRoleNotFound,
		"bad_request:role_is_used":   ErrRoleIsUsed,
	}

	// Parse errors
	if err, ok := errMap[errMessage]; ok {
		return err
	}

	return fmt.Errorf("unexpected response: status code: %d, message: %s", res.StatusCode(), errMessage)
}
