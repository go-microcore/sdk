package adapter

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"

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

func (a *adapter) Auth(ctx context.Context, data AuthData) (*AuthResult, error) {
	// Encode body json
	body, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error parsing request body: %v", err)
	}

	// Encode body
	encBody, err := a.encrypt(body)
	if err != nil {
		return nil, fmt.Errorf("error encode body: %v", err)
	}

	// Send service request
	res, err := a.httpClientManager.Request(
		a.authServiceEndpoint+"/auth/",
		client.WithRequestMethod(http.MethodPost),
		client.WithRequestBody(encBody),
		client.WithRequestContext(ctx),
	)
	if err != nil {
		return nil, fmt.Errorf("service %s unavailable: %v", a.authServiceEndpoint, err)
	}

	// Check success status code
	if res.StatusCode() == 200 {
		var response AuthResult
		if err := json.Unmarshal(res.Body(), &response); err != nil {
			return nil, fmt.Errorf("error parsing response body: %v", err)
		}
		return &response, nil
	}

	// Response message
	errMessage := string(res.Body())

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
	// Encode body json
	body, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error parsing request body: %v", err)
	}

	// Encode body
	encBody, err := a.encrypt(body)
	if err != nil {
		return nil, fmt.Errorf("error encode body: %v", err)
	}

	// Send service request
	res, err := a.httpClientManager.Request(
		a.authServiceEndpoint+"/auth/2fa",
		client.WithRequestMethod(http.MethodPost),
		client.WithRequestBody(encBody),
		client.WithRequestContext(ctx),
	)
	if err != nil {
		return nil, fmt.Errorf("service %s unavailable: %v", a.authServiceEndpoint, err)
	}

	// Check success status code
	if res.StatusCode() == 200 {
		var response Auth2faResult
		if err := json.Unmarshal(res.Body(), &response); err != nil {
			return nil, fmt.Errorf("error parsing response body: %v", err)
		}
		return &response, nil
	}

	// Response message
	errMessage := string(res.Body())

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
	// Encode body json
	body, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error parsing request body: %v", err)
	}

	// Send service request
	res, err := a.httpClientManager.Request(
		a.authServiceEndpoint+"/auth/token/renew",
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

func (a *adapter) TokenValidate(ctx context.Context, data TokenValidateData) (*TokenValidateResult, error) {
	// Encode body json
	body, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error parsing request body: %v", err)
	}

	// Send service request
	res, err := a.httpClientManager.Request(
		a.authServiceEndpoint+"/auth/token/validate",
		client.WithRequestMethod(http.MethodPost),
		client.WithRequestBody(body),
		client.WithRequestContext(ctx),
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

func (a *adapter) UserLogout(ctx context.Context, authToken string) error {
	// Send service request
	res, err := a.httpClientManager.Request(
		a.authServiceEndpoint+"/auth/user/logout/",
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

	return fmt.Errorf("unexpected response: status code: %d", res.StatusCode())
}

func (a *adapter) UserLogoutAll(ctx context.Context, authToken string) error {
	// Send service request
	res, err := a.httpClientManager.Request(
		a.authServiceEndpoint+"/auth/user/logout/all",
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

	return fmt.Errorf("unexpected response: status code: %d", res.StatusCode())
}

func (a *adapter) UserLogoutDevice(ctx context.Context, authToken string, data LogoutDeviceData) error {
	// Encode body json
	body, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("error parsing request body: %v", err)
	}

	// Send service request
	res, err := a.httpClientManager.Request(
		a.authServiceEndpoint+"/auth/user/logout/device",
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

	return fmt.Errorf("unexpected response: status code: %d", res.StatusCode())
}

func (a *adapter) UserDevices(ctx context.Context, authToken string) ([]DeviceResult, error) {
	// Send service request
	res, err := a.httpClientManager.Request(
		a.authServiceEndpoint+"/auth/user/devices",
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

	return nil, fmt.Errorf("unexpected response: status code: %d, message: %s", res.StatusCode(), string(res.Body()))
}

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
