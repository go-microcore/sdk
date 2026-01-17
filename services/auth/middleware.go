package auth

import (
	"encoding/json"
	"strings"

	"go.microcore.dev/framework/errors"
	"go.microcore.dev/framework/transport/http"
	"go.microcore.dev/framework/transport/http/client"
	"go.microcore.dev/framework/transport/http/server"
)

var (
	ErrAuthInsufficientPermissions = errors.New(errors.ErrForbidden, "insufficient_permissions")
	ErrAuth2faRequired             = errors.New(errors.ErrUnauthorized, "2fa_required")
	ErrAuthInvalidToken            = errors.New(errors.ErrUnauthorized, "invalid_token")
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
	Id       string   `json:"id"`
	Device   string   `json:"device"`
	User     uint     `json:"user"`
	Roles    []string `json:"roles"`
	Mfa      bool     `json:"mfa"`
	Expires  int64    `json:"expires"`
	Issued   int64    `json:"issued"`
	Issuer   string   `json:"issuer"`
	Audience []string `json:"audience"`
}
type tokenAuthorizeHttpAuthResponse struct {
	Mfa bool `json:"mfa"`
}

type MiddlewareConfig struct {
	AuthServiceEndpoint string
	HttpClientManager   client.Manager
}

func NewMiddleware(config *MiddlewareConfig) *middleware {
	return &middleware{
		authServiceEndpoint: config.AuthServiceEndpoint,
		httpClientManager:   config.HttpClientManager,
	}
}

type middleware struct {
	authServiceEndpoint string
	httpClientManager   client.Manager
}

func (m *middleware) Auth() func(server.RequestHandler) server.RequestHandler {
	return func(handler server.RequestHandler) server.RequestHandler {
		return func(c *server.RequestContext) {
			// Build url
			var url strings.Builder
			url.WriteString(m.authServiceEndpoint)
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
				c.WriteError(errors.ErrServiceUnavailable)
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
				c.WriteError(errors.ErrServiceUnavailable)
				return
			}

			// Check status code
			switch res.StatusCode() {
			case 200:
				// Parse response
				var response tokenAuthorizeHttpResponse
				if err := json.Unmarshal(res.Body(), &response); err != nil {
					c.WriteError(errors.ErrServiceUnavailable)
					return
				}

				// Set data to ctx
				c.SetUserValue("device", response.Token.Device)
				c.SetUserValue("user", response.Token.User)
				c.SetUserValue("roles", response.Token.Roles)
				c.SetUserValue("mfa_value", response.Token.Mfa)
				c.SetUserValue("mfa_validation", response.Auth.Mfa)

				// Check two factor
				if response.Auth.Mfa && response.Token.Mfa {
					c.WriteError(ErrAuth2faRequired)
					return
				}

				handler(c)
			case 400:
				c.WriteError(ErrAuthInvalidToken)
			case 403:
				c.WriteError(ErrAuthInsufficientPermissions)
			default:
				c.WriteError(errors.ErrServiceUnavailable)
			}
		}
	}
}
