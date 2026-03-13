package auth

import (
	"encoding/json"
	"fmt"
	"strings"

	"go.microcore.dev/framework/transport"
	"go.microcore.dev/framework/transport/http"
	"go.microcore.dev/framework/transport/http/client"
	"go.microcore.dev/framework/transport/http/server"
)

var (
	ErrAuthInsufficientPermissions = transport.NewError(
		transport.ErrForbidden,
		"insufficient permissions",
		"INSUFFICIENT_PERMISSIONS",
	)

	ErrAuth2faRequired = transport.NewError(
		transport.ErrUnauthorized,
		"2fa required",
		"2FA_REQUIRED",
	)

	ErrAuthInvalidToken = transport.NewError(
		transport.ErrUnauthorized,
		"invalid token",
		"INVALID_TOKEN",
	)
)

type tokenAuthorizeHttpRequest struct {
	Path   string `json:"path"`
	Method string `json:"method"`
}

type tokenAuthorizeHttpResponse struct {
	Token tokenAuthorizeHttpDataResponse `json:"token"`
	Auth  tokenAuthorizeHttpAuthResponse `json:"auth"`
}
type tokenAuthorizeHttpDataResponse struct {
	ID       string   `json:"id"`
	Device   string   `json:"device"`
	User     uint     `json:"user"`
	Roles    []string `json:"roles"`
	Mfa      bool     `json:"mfa"`
	Expires  *int64   `json:"expires"`
	Issued   int64    `json:"issued"`
	Issuer   string   `json:"issuer"`
	Audience []string `json:"audience"`
}
type tokenAuthorizeHttpAuthResponse struct {
	Mfa bool `json:"mfa"`
}

type MiddlewareConfig struct {
	AuthServiceServer string
	HttpClientManager client.Manager
}

func NewMiddleware(config *MiddlewareConfig) *Middleware {
	return &Middleware{
		authServiceServer: config.AuthServiceServer,
		httpClientManager: config.HttpClientManager,
	}
}

type Middleware struct {
	authServiceServer string
	httpClientManager client.Manager
}

func (m *Middleware) Handler(handler server.RequestHandler) server.RequestHandler {
	return func(c *server.RequestContext) {
		// Build url
		var url strings.Builder
		url.WriteString(m.authServiceServer)
		url.WriteString("/auth/tokens/authorize/http")

		// Get auth token
		token, err := c.GetBearerToken()
		if err != nil {
			c.WriteError(ErrAuthInvalidToken)
			return
		}

		// Encode body json
		body, err := json.Marshal(
			tokenAuthorizeHttpRequest{
				Path:   string(c.Path()),
				Method: string(c.Method()),
			},
		)
		if err != nil {
			c.WriteError(fmt.Errorf("marshal request: %w", err))
			return
		}

		// Authorize HTTP JWT token
		res, err := m.httpClientManager.Request(
			url.String(),
			client.WithRequestMethod(http.MethodPost),
			client.WithRequestContext(c.GetContext()),
			client.WithRequestBody(body),
			client.WithRequestHeaders(
				client.NewRequestHeader("Authorization", "Bearer "+token),
			),
		)
		if err != nil {
			c.WriteError(fmt.Errorf("authorize http roles request: %w", err))
			return
		}

		// Check status code
		switch res.StatusCode() {
		case 200:
			// Parse response
			var response tokenAuthorizeHttpResponse
			if err := json.Unmarshal(res.Body(), &response); err != nil {
				c.WriteError(fmt.Errorf("unmarshal response: %w", err))
				return
			}

			// Check two factor
			if response.Auth.Mfa && response.Token.Mfa {
				c.WriteError(ErrAuth2faRequired)
				return
			}

			// Set data to ctx
			c.SetUserValue("device", response.Token.Device)
			c.SetUserValue("user", response.Token.User)
			c.SetUserValue("roles", response.Token.Roles)
			c.SetUserValue("mfa_value", response.Token.Mfa)
			c.SetUserValue("mfa_validation", response.Auth.Mfa)

			handler(c)
		case 400:
			c.WriteError(ErrAuthInvalidToken)
		case 401:
			c.WriteError(ErrAuthInvalidToken)
		case 403:
			c.WriteError(ErrAuthInsufficientPermissions)
		default:
			c.WriteError(transport.ErrServiceUnavailable)
		}
	}
}
