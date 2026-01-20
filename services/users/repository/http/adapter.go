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

func (a *adapter) TwoFASettings(ctx context.Context, authToken string, data TwoFASettingsData) (*TwoFASettingsResult, error) {
	// Build url
	var url strings.Builder
	url.WriteString(a.usersServiceEndpoint)
	url.WriteString("/users/2fa/settings/")
	
	// Encode body json
	body, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error parsing request body: %v", err)
	}

	// Send service request
	res, err := a.httpClientManager.Request(
		url.String(),
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
		var response TwoFASettingsResult
		if err := json.Unmarshal(res.Body(), &response); err != nil {
			return nil, fmt.Errorf("error parsing response body: %v", err)
		}
		return &response, nil
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
	// Build url
	var url strings.Builder
	url.WriteString(a.usersServiceEndpoint)
	url.WriteString("/users/2fa/settings/enable")
	
	// Encode body json
	body, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("error parsing request body: %v", err)
	}

	// Send service request
	res, err := a.httpClientManager.Request(
		url.String(),
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
	// Build url
	var url strings.Builder
	url.WriteString(a.usersServiceEndpoint)
	url.WriteString("/users/2fa/settings/disable")
	
	// Encode body json
	body, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("error parsing request body: %v", err)
	}

	// Send service request
	res, err := a.httpClientManager.Request(
		url.String(),
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

func (a *adapter) TwoFAValidate(ctx context.Context, authToken string, data TwoFAValidateData) (*TwoFAValidateResult, error) {
	// Build url
	var url strings.Builder
	url.WriteString(a.usersServiceEndpoint)
	url.WriteString("/users/2fa/validate")
	
	// Encode body json
	body, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error parsing request body: %v", err)
	}

	// Send service request
	res, err := a.httpClientManager.Request(
		url.String(),
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
		var response TwoFAValidateResult
		if err := json.Unmarshal(res.Body(), &response); err != nil {
			return nil, fmt.Errorf("error parsing response body: %v", err)
		}
		return &response, nil
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

func (a *adapter) Signin(ctx context.Context, data SigninData) (*SigninResult, error) {
	// Build url
	var url strings.Builder
	url.WriteString(a.usersServiceEndpoint)
	url.WriteString("/users/signin")
	
	// Encode body json
	body, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error parsing request body: %v", err)
	}

	// Send service request
	res, err := a.httpClientManager.Request(
		url.String(),
		client.WithRequestMethod(http.MethodPost),
		client.WithRequestBody(body),
		client.WithRequestContext(ctx),
	)
	if err != nil {
		return nil, fmt.Errorf("service %s unavailable: %v", a.usersServiceEndpoint, err)
	}

	// Check success status code
	if res.StatusCode() == 200 {
		var response SigninResult
		if err := json.Unmarshal(res.Body(), &response); err != nil {
			return nil, fmt.Errorf("error parsing response body: %v", err)
		}
		return &response, nil
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
	// Build url
	var url strings.Builder
	url.WriteString(a.usersServiceEndpoint)
	url.WriteString("/users/signup")
	
	// Encode body json
	body, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error parsing request body: %v", err)
	}

	// Send service request
	res, err := a.httpClientManager.Request(
		url.String(),
		client.WithRequestMethod(http.MethodPost),
		client.WithRequestBody(body),
		client.WithRequestContext(ctx),
	)
	if err != nil {
		return nil, fmt.Errorf("service %s unavailable: %v", a.usersServiceEndpoint, err)
	}

	// Check success status code
	if res.StatusCode() == 201 {
		var response SignupResult
		if err := json.Unmarshal(res.Body(), &response); err != nil {
			return nil, fmt.Errorf("error parsing response body: %v", err)
		}
		return &response, nil
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
	// Build url
	var url strings.Builder
	url.WriteString(a.usersServiceEndpoint)
	url.WriteString("/users/profile")
	
	// Send service request
	res, err := a.httpClientManager.Request(
		url.String(),
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
		var response ProfileResult
		if err := json.Unmarshal(res.Body(), &response); err != nil {
			return nil, fmt.Errorf("error parsing response body: %v", err)
		}
		return &response, nil
	}

	// Response message
	errMessage := string(res.Body())

	return nil, fmt.Errorf("unexpected response: status code: %d, message: %s", res.StatusCode(), errMessage)
}

func (a *adapter) FilterUsers(ctx context.Context, authToken string, data FilterUsersData) ([]FilterUsersResult, error) {
	// Build url
	var url strings.Builder
	url.WriteString(a.usersServiceEndpoint)
	url.WriteString("/users/filter")
	
	// Encode body json
	body, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error parsing request body: %v", err)
	}

	// Send service request
	res, err := a.httpClientManager.Request(
		url.String(),
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
		var response []FilterUsersResult
		if err := json.Unmarshal(res.Body(), &response); err != nil {
			return nil, fmt.Errorf("error parsing response body: %v", err)
		}
		return response, nil
	}

	// Response message
	errMessage := string(res.Body())

	return nil, fmt.Errorf("unexpected response: status code: %d, message: %s", res.StatusCode(), errMessage)
}

func (a *adapter) UpdateUser(ctx context.Context, authToken string, id string, data UpdateUserData) error {
	// Build url
	var url strings.Builder
	url.WriteString(a.usersServiceEndpoint)
	url.WriteString("/users/")
	url.WriteString(id)
	
	// Encode body json
	body, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("error parsing request body: %v", err)
	}

	// Send service request
	res, err := a.httpClientManager.Request(
		url.String(),
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
	// Build url
	var url strings.Builder
	url.WriteString(a.usersServiceEndpoint)
	url.WriteString("/users/")
	url.WriteString(strconv.FormatUint(uint64(id), 10))

	// Send service request
	res, err := a.httpClientManager.Request(
		url.String(),
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

func (a *adapter) CreateUser(ctx context.Context, authToken string, data CreateUserData) (*CreateUserResult, error) {
	// Build url
	var url strings.Builder
	url.WriteString(a.usersServiceEndpoint)
	url.WriteString("/users/")
	
	// Encode body json
	body, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error parsing request body: %v", err)
	}

	// Send service request
	res, err := a.httpClientManager.Request(
		url.String(),
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
		var response CreateUserResult
		if err := json.Unmarshal(res.Body(), &response); err != nil {
			return nil, fmt.Errorf("error parsing response body: %v", err)
		}
		return &response, nil
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
