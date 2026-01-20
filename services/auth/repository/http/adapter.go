package adapter

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"

	"go.microcore.dev/framework/errors"
	"go.microcore.dev/framework/transport/http"
	"go.microcore.dev/framework/transport/http/client"
)

type Config struct {
	HttpClientManager   client.Manager
	AuthServiceEndpoint string
	AuthKey             string
}

func New(config *Config) Interface {
	return &adapter{
		config.HttpClientManager,
		config.AuthServiceEndpoint,
		config.AuthKey,
	}
}

type adapter struct {
	httpClientManager   client.Manager
	authServiceEndpoint string
	authKey             string
}

// Devices

func (a *adapter) GetDevices(ctx context.Context, authToken string) ([]DeviceResult, error) {
	// Build url
	var url strings.Builder
	url.WriteString(a.authServiceEndpoint)
	url.WriteString("/auth/devices")

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
		return nil, fmt.Errorf("service %s unavailable: %v", a.authServiceEndpoint, err)
	}

	// Check success status code
	if res.StatusCode() == 200 {
		var response []DeviceResult
		if err := json.Unmarshal(res.Body(), &response); err != nil {
			return nil, fmt.Errorf("error parsing response body: %v", err)
		}
		return response, nil
	}

	// Response message
	errMessage := string(res.Body())

	return nil, fmt.Errorf("unexpected response: status code: %d, message: %s", res.StatusCode(), errMessage)
}

// Logout

func (a *adapter) Logout(ctx context.Context, authToken string) error {
	// Build url
	var url strings.Builder
	url.WriteString(a.authServiceEndpoint)
	url.WriteString("/auth/logout/")

	// Send service request
	res, err := a.httpClientManager.Request(
		url.String(),
		client.WithRequestMethod(http.MethodPost),
		client.WithRequestContext(ctx),
		client.WithRequestHeaders(
			client.NewRequestHeader("Authorization", "Bearer "+authToken),
		),
	)
	if err != nil {
		return fmt.Errorf("service %s unavailable: %v", a.authServiceEndpoint, err)
	}

	// Check success status code
	if res.StatusCode() == 204 {
		return nil
	}

	// Response message
	errMessage := string(res.Body())

	return fmt.Errorf("unexpected response: status code: %d, message: %s", res.StatusCode(), errMessage)
}

func (a *adapter) LogoutAll(ctx context.Context, authToken string) error {
	// Build url
	var url strings.Builder
	url.WriteString(a.authServiceEndpoint)
	url.WriteString("/auth/logout/all")

	// Send service request
	res, err := a.httpClientManager.Request(
		url.String(),
		client.WithRequestMethod(http.MethodPost),
		client.WithRequestContext(ctx),
		client.WithRequestHeaders(
			client.NewRequestHeader("Authorization", "Bearer "+authToken),
		),
	)
	if err != nil {
		return fmt.Errorf("service %s unavailable: %v", a.authServiceEndpoint, err)
	}

	// Check success status code
	if res.StatusCode() == 204 {
		return nil
	}

	// Response message
	errMessage := string(res.Body())

	return fmt.Errorf("unexpected response: status code: %d, message: %s", res.StatusCode(), errMessage)
}

func (a *adapter) LogoutDevice(ctx context.Context, authToken string, data LogoutDeviceData) error {
	// Build url
	var url strings.Builder
	url.WriteString(a.authServiceEndpoint)
	url.WriteString("/auth/logout/device")

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
		return fmt.Errorf("service %s unavailable: %v", a.authServiceEndpoint, err)
	}

	// Check success status code
	if res.StatusCode() == 204 {
		return nil
	}

	// Response message
	errMessage := string(res.Body())

	// Errors map
	var errMap = map[string]error{
		"bad_request:invalid_device": ErrInvalidDevice,
	}

	// Parse errors
	if err, ok := errMap[errMessage]; ok {
		return err
	}

	return fmt.Errorf("unexpected response: status code: %d, message: %s", res.StatusCode(), errMessage)
}

// Roles

func (a *adapter) CreateRole(ctx context.Context, authToken string, data CreateRoleData) (*CreateRoleResult, error) {
	// Build url
	var url strings.Builder
	url.WriteString(a.authServiceEndpoint)
	url.WriteString("/auth/roles/")

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
		return nil, fmt.Errorf("service %s unavailable: %v", a.authServiceEndpoint, err)
	}

	// Check success status code
	if res.StatusCode() == 201 {
		var response CreateRoleResult
		if err := json.Unmarshal(res.Body(), &response); err != nil {
			return nil, fmt.Errorf("error parsing response body: %v", err)
		}
		return &response, nil
	}

	// Response message
	errMessage := string(res.Body())

	// Errors map
	var errMap = map[string]error{
		"bad_request:invalid_role_id":          ErrInvalidRoleId,
		"bad_request:invalid_role_name":        ErrInvalidRoleName,
		"bad_request:invalid_role_description": ErrInvalidRoleDescription,
		"bad_request:role_exist_id":            ErrRoleExistId,
	}

	// Parse errors
	if err, ok := errMap[errMessage]; ok {
		return nil, err
	}

	return nil, fmt.Errorf("unexpected response: status code: %d, message: %s", res.StatusCode(), errMessage)
}

func (a *adapter) FilterRoles(ctx context.Context, authToken string, data FilterRolesData) ([]FilterRolesResult, error) {
	// Build url
	var url strings.Builder
	url.WriteString(a.authServiceEndpoint)
	url.WriteString("/auth/roles/filter")

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
		return nil, fmt.Errorf("service %s unavailable: %v", a.authServiceEndpoint, err)
	}

	// Check success status code
	if res.StatusCode() == 200 {
		var response []FilterRolesResult
		if err := json.Unmarshal(res.Body(), &response); err != nil {
			return nil, fmt.Errorf("error parsing response body: %v", err)
		}
		return response, nil
	}

	// Response message
	errMessage := string(res.Body())

	return nil, fmt.Errorf("unexpected response: status code: %d, message: %s", res.StatusCode(), errMessage)
}

func (a *adapter) UpdateRole(ctx context.Context, authToken string, id string, data UpdateRoleData) error {
	// Build url
	var url strings.Builder
	url.WriteString(a.authServiceEndpoint)
	url.WriteString("/auth/roles/")
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
		return fmt.Errorf("service %s unavailable: %v", a.authServiceEndpoint, err)
	}

	// Check success status code
	if res.StatusCode() == 204 {
		return nil
	}

	// Response message
	errMessage := string(res.Body())

	// Errors map
	var errMap = map[string]error{
		"bad_request:invalid_role_id":          ErrInvalidRoleId,
		"bad_request:invalid_role_name":        ErrInvalidRoleName,
		"bad_request:invalid_role_description": ErrInvalidRoleDescription,
		"bad_request:role_not_found":           ErrRoleNotFound,
	}

	// Parse errors
	if err, ok := errMap[errMessage]; ok {
		return err
	}

	return fmt.Errorf("unexpected response: status code: %d, message: %s", res.StatusCode(), errMessage)
}

func (a *adapter) DeleteRole(ctx context.Context, authToken string, id string) error {
	// Build url
	var url strings.Builder
	url.WriteString(a.authServiceEndpoint)
	url.WriteString("/auth/roles/")
	url.WriteString(id)

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
		return fmt.Errorf("service %s unavailable: %v", a.authServiceEndpoint, err)
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
	}

	// Parse errors
	if err, ok := errMap[errMessage]; ok {
		return err
	}

	return fmt.Errorf("unexpected response: status code: %d, message: %s", res.StatusCode(), errMessage)
}

// Rules (HTTP)

func (a *adapter) CreateHttpRule(ctx context.Context, authToken string, data CreateHttpRuleData) (*CreateHttpRuleResult, error) {
	// Build url
	var url strings.Builder
	url.WriteString(a.authServiceEndpoint)
	url.WriteString("/auth/rules/http/")

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
		return nil, fmt.Errorf("service %s unavailable: %v", a.authServiceEndpoint, err)
	}

	// Check success status code
	if res.StatusCode() == 201 {
		var response CreateHttpRuleResult
		if err := json.Unmarshal(res.Body(), &response); err != nil {
			return nil, fmt.Errorf("error parsing response body: %v", err)
		}
		return &response, nil
	}

	// Response message
	errMessage := string(res.Body())

	// Errors map
	var errMap = map[string]error{
		"bad_request:invalid_role_id": ErrInvalidRoleId,
		"bad_request:invalid_path":    ErrInvalidPath,
		"bad_request:invalid_methods": ErrInvalidMethods,
		"bad_request:rule_exist":      ErrRuleExist,
	}

	// Parse errors
	if err, ok := errMap[errMessage]; ok {
		return nil, err
	}

	return nil, fmt.Errorf("unexpected response: status code: %d, message: %s", res.StatusCode(), errMessage)
}

func (a *adapter) FilterHttpRules(ctx context.Context, authToken string, data FilterHttpRulesData) ([]FilterHttpRulesResult, error) {
	// Build url
	var url strings.Builder
	url.WriteString(a.authServiceEndpoint)
	url.WriteString("/auth/rules/http/filter")

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
		return nil, fmt.Errorf("service %s unavailable: %v", a.authServiceEndpoint, err)
	}

	// Check success status code
	if res.StatusCode() == 200 {
		var response []FilterHttpRulesResult
		if err := json.Unmarshal(res.Body(), &response); err != nil {
			return nil, fmt.Errorf("error parsing response body: %v", err)
		}
		return response, nil
	}
	
	// Response message
	errMessage := string(res.Body())

	return nil, fmt.Errorf("unexpected response: status code: %d, message: %s", res.StatusCode(), errMessage)
}

func (a *adapter) UpdateHttpRule(ctx context.Context, authToken string, id uint, data UpdateHttpRuleData) error {
	// Build url
	var url strings.Builder
	url.WriteString(a.authServiceEndpoint)
	url.WriteString("/auth/rules/http/")
	url.WriteString(strconv.FormatUint(uint64(id), 10))

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
		return fmt.Errorf("service %s unavailable: %v", a.authServiceEndpoint, err)
	}

	// Check success status code
	if res.StatusCode() == 204 {
		return nil
	}

	// Response message
	errMessage := string(res.Body())

	// Errors map
	var errMap = map[string]error{
		"bad_request:invalid_role_id": ErrInvalidRoleId,
		"bad_request:invalid_path":    ErrInvalidPath,
		"bad_request:invalid_methods": ErrInvalidMethods,
		"bad_request:invalid_mfa":     ErrInvalidMfa,
		"bad_request:rule_not_found":  ErrRuleNotFound,
	}

	// Parse errors
	if err, ok := errMap[errMessage]; ok {
		return err
	}

	return fmt.Errorf("unexpected response: status code: %d, message: %s", res.StatusCode(), errMessage)
}

func (a *adapter) DeleteHttpRule(ctx context.Context, authToken string, id uint) error {
	// Build url
	var url strings.Builder
	url.WriteString(a.authServiceEndpoint)
	url.WriteString("/auth/rules/http/")
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
		return fmt.Errorf("service %s unavailable: %v", a.authServiceEndpoint, err)
	}

	// Check success status code
	if res.StatusCode() == 204 {
		return nil
	}

	// Response message
	errMessage := string(res.Body())

	// Errors map
	var errMap = map[string]error{
		"bad_request:rule_not_found": ErrRuleNotFound,
	}

	// Parse errors
	if err, ok := errMap[errMessage]; ok {
		return err
	}

	return fmt.Errorf("unexpected response: status code: %d, message: %s", res.StatusCode(), errMessage)
}

// Tokens

func (a *adapter) Auth(ctx context.Context, data AuthData) (*AuthResult, error) {
	// Build url
	var url strings.Builder
	url.WriteString(a.authServiceEndpoint)
	url.WriteString("/auth/tokens/")

	// Encode body json
	requestbody, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error parsing request body: %v", err)
	}

	// Encode request body
	encRequestBody, err := a.encrypt(requestbody)
	if err != nil {
		return nil, fmt.Errorf("error encode body: %v", err)
	}

	// Send service request
	res, err := a.httpClientManager.Request(
		url.String(),
		client.WithRequestMethod(http.MethodPost),
		client.WithRequestBody(encRequestBody),
		client.WithRequestContext(ctx),
	)
	if err != nil {
		return nil, fmt.Errorf("service %s unavailable: %v", a.authServiceEndpoint, err)
	}

	// Decode response body
	decResponseBody, err := a.decrypt(res.Body())
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt auth service response: %v", err)
	}

	// Check success status code
	if res.StatusCode() == 200 {
		var response AuthResult
		if err := json.Unmarshal(decResponseBody, &response); err != nil {
			return nil, fmt.Errorf("error parsing response body: %v", err)
		}
		return &response, nil
	}

	// Response message
	errMessage := string(decResponseBody)

	// Errors map
	var errMap = map[string]error{
		"bad_request": errors.ErrBadRequest,
	}

	// Parse errors
	if err, ok := errMap[errMessage]; ok {
		return nil, err
	}

	return nil, fmt.Errorf("unexpected response: status code: %d, message: %s", res.StatusCode(), errMessage)
}

func (a *adapter) Auth2fa(ctx context.Context, data Auth2faData) (*Auth2faResult, error) {
	// Build url
	var url strings.Builder
	url.WriteString(a.authServiceEndpoint)
	url.WriteString("/auth/tokens/2fa")

	// Encode body json
	requestbody, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error parsing request body: %v", err)
	}

	// Encode request body
	encRequestBody, err := a.encrypt(requestbody)
	if err != nil {
		return nil, fmt.Errorf("error encode body: %v", err)
	}

	// Send service request
	res, err := a.httpClientManager.Request(
		url.String(),
		client.WithRequestMethod(http.MethodPost),
		client.WithRequestBody(encRequestBody),
		client.WithRequestContext(ctx),
	)
	if err != nil {
		return nil, fmt.Errorf("service %s unavailable: %v", a.authServiceEndpoint, err)
	}

	// Decode response body
	decResponseBody, err := a.decrypt(res.Body())
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt auth service response: %v", err)
	}

	// Check success status code
	if res.StatusCode() == 200 {
		var response Auth2faResult
		if err := json.Unmarshal(decResponseBody, &response); err != nil {
			return nil, fmt.Errorf("error parsing response body: %v", err)
		}
		return &response, nil
	}

	// Response message
	errMessage := string(decResponseBody)

	// Errors map
	var errMap = map[string]error{
		"bad_request": errors.ErrBadRequest,
	}

	// Parse errors
	if err, ok := errMap[errMessage]; ok {
		return nil, err
	}

	return nil, fmt.Errorf("unexpected response: status code: %d, message: %s", res.StatusCode(), errMessage)
}

func (a *adapter) TokenRenew(ctx context.Context, data TokenRenewData) (*TokenRenewResult, error) {
	// Build url
	var url strings.Builder
	url.WriteString(a.authServiceEndpoint)
	url.WriteString("/auth/tokens/renew")

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
		return nil, fmt.Errorf("service %s unavailable: %v", a.authServiceEndpoint, err)
	}

	// Check success status code
	if res.StatusCode() == 200 {
		var response TokenRenewResult
		if err := json.Unmarshal(res.Body(), &response); err != nil {
			return nil, fmt.Errorf("error parsing response body: %v", err)
		}
		return &response, nil
	}

	// Response message
	errMessage := string(res.Body())

	// Errors map
	var errMap = map[string]error{
		"bad_request:invalid_token":      ErrInvalidToken,
		"bad_request:token_already_used": ErrTokenAlreadyUsed,
	}

	// Parse errors
	if err, ok := errMap[errMessage]; ok {
		return nil, err
	}

	return nil, fmt.Errorf("unexpected response: status code: %d, message: %s", res.StatusCode(), errMessage)
}

func (a *adapter) TokenValidate(ctx context.Context, authToken string) (*TokenValidateResult, error) {
	// Build url
	var url strings.Builder
	url.WriteString(a.authServiceEndpoint)
	url.WriteString("/auth/tokens/validate")

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
		return nil, fmt.Errorf("service %s unavailable: %v", a.authServiceEndpoint, err)
	}

	// Check success status code
	if res.StatusCode() == 200 {
		var response TokenValidateResult
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
	}

	// Parse errors
	if err, ok := errMap[errMessage]; ok {
		return nil, err
	}

	return nil, fmt.Errorf("unexpected response: status code: %d, message: %s", res.StatusCode(), errMessage)
}

func (a *adapter) TokenAuthorizeHttp(ctx context.Context, authToken string, data TokenAuthorizeHttpData) (*TokenAuthorizeHttpResult, error) {
	// Build url
	var url strings.Builder
	url.WriteString(a.authServiceEndpoint)
	url.WriteString("/auth/tokens/authorize/http")

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
		return nil, fmt.Errorf("service %s unavailable: %v", a.authServiceEndpoint, err)
	}

	// Check success status code
	if res.StatusCode() == 200 {
		var response TokenAuthorizeHttpResult
		if err := json.Unmarshal(res.Body(), &response); err != nil {
			return nil, fmt.Errorf("error parsing response body: %v", err)
		}
		return &response, nil
	}

	// Response message
	errMessage := string(res.Body())

	// Errors map
	var errMap = map[string]error{
		"bad_request:invalid_token":          ErrInvalidToken,
		"bad_request:invalid_path":           ErrInvalidPath,
		"bad_request:invalid_methods":        ErrInvalidMethods,
		"forbidden:insufficient_permissions": ErrInsufficientPermissions,
	}

	// Parse errors
	if err, ok := errMap[errMessage]; ok {
		return nil, err
	}

	return nil, fmt.Errorf("unexpected response: status code: %d, message: %s", res.StatusCode(), errMessage)
}

// Static access tokens

func (a *adapter) CreateStaticAccessToken(ctx context.Context, authToken string, data CreateStaticAccessTokenData) (*CreateStaticAccessTokenResult, error) {
	// Build url
	var url strings.Builder
	url.WriteString(a.authServiceEndpoint)
	url.WriteString("/auth/tokens/static/")

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
		return nil, fmt.Errorf("service %s unavailable: %v", a.authServiceEndpoint, err)
	}

	// Check success status code
	if res.StatusCode() == 201 {
		var response CreateStaticAccessTokenResult
		if err := json.Unmarshal(res.Body(), &response); err != nil {
			return nil, fmt.Errorf("error parsing response body: %v", err)
		}
		return &response, nil
	}

	// Response message
	errMessage := string(res.Body())

	// Errors map
	var errMap = map[string]error{
		"bad_request:invalid_id":          ErrInvalidId,
		"bad_request:invalid_roles":       ErrInvalidRoles,
		"bad_request:invalid_description": ErrInvalidDescription,
		"bad_request:static_token_exist":  ErrStaticTokenExist,
	}

	// Parse errors
	if err, ok := errMap[errMessage]; ok {
		return nil, err
	}

	return nil, fmt.Errorf("unexpected response: status code: %d, message: %s", res.StatusCode(), errMessage)
}

func (a *adapter) FilterStaticAccessTokens(ctx context.Context, authToken string, data FilterStaticAccessTokenData) ([]FilterStaticAccessTokenResult, error) {
	// Build url
	var url strings.Builder
	url.WriteString(a.authServiceEndpoint)
	url.WriteString("/auth/tokens/static/filter")

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
		return nil, fmt.Errorf("service %s unavailable: %v", a.authServiceEndpoint, err)
	}

	// Check success status code
	if res.StatusCode() == 200 {
		var response []FilterStaticAccessTokenResult
		if err := json.Unmarshal(res.Body(), &response); err != nil {
			return nil, fmt.Errorf("error parsing response body: %v", err)
		}
		return response, nil
	}

	// Response message
	errMessage := string(res.Body())

	return nil, fmt.Errorf("unexpected response: status code: %d, message: %s", res.StatusCode(), errMessage)
}

func (a *adapter) DeleteStaticAccessToken(ctx context.Context, authToken string, id string) error {
	// Build url
	var url strings.Builder
	url.WriteString(a.authServiceEndpoint)
	url.WriteString("/auth/tokens/static/")
	url.WriteString(id)

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
		return fmt.Errorf("service %s unavailable: %v", a.authServiceEndpoint, err)
	}

	// Check success status code
	if res.StatusCode() == 204 {
		return nil
	}

	// Response message
	errMessage := string(res.Body())

	// Errors map
	var errMap = map[string]error{
		"bad_request:static_token_not_found": ErrStaticTokenNotFound,
	}

	// Parse errors
	if err, ok := errMap[errMessage]; ok {
		return err
	}

	return fmt.Errorf("unexpected response: status code: %d, message: %s", res.StatusCode(), errMessage)
}

// Helper for encrypt auth request data
func (a *adapter) encrypt(data []byte) ([]byte, error) {
	block, err := aes.NewCipher([]byte(a.authKey))
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	ciphertext := aesGCM.Seal(nonce, nonce, data, nil)
	return ciphertext, nil
}

// Helper for decrypt auth response data
func (a *adapter) decrypt(data []byte) ([]byte, error) {
	block, err := aes.NewCipher([]byte(a.authKey))
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := aesGCM.NonceSize()
	if len(data) < nonceSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	nonce, data := data[:nonceSize], data[nonceSize:]
	res, err := aesGCM.Open(nil, nonce, data, nil)
	if err != nil {
		return nil, err
	}

	return res, nil
}
