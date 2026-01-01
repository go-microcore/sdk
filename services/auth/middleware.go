package auth

import (
	"encoding/json"
	"slices"

	"go.microcore.dev/framework/errors"
	"go.microcore.dev/framework/transport/http"
	"go.microcore.dev/framework/transport/http/client"
	"go.microcore.dev/framework/transport/http/server"
)

type MiddlewareConfig struct {
	AuthServiceEndpoint string
	HttpClientManager   client.Manager
}

func NewMiddleware(config *MiddlewareConfig) Middleware {
	return &middleware{
		authServiceEndpoint: config.AuthServiceEndpoint,
		httpClientManager:   config.HttpClientManager,
	}
}

type Middleware interface {
	Auth(options ...AuthOption) func(server.RequestHandler) server.RequestHandler
}

type middleware struct {
	authServiceEndpoint string
	httpClientManager   client.Manager
}

type tokenValidateRequest struct {
	AccessToken string `json:"access_token"`
}

type tokenValidateResult struct {
	Id       string   `json:"id"`
	Device   string   `json:"device"`
	User     uint     `json:"user"`
	Role     string   `json:"role"`
	Mfa      bool     `json:"mfa"`
	Expires  int64    `json:"expires"`
	Issued   int64    `json:"issued"`
	Issuer   string   `json:"issuer"`
	Audience []string `json:"audience"`
}

func (m *middleware) Auth(options ...AuthOption) func(server.RequestHandler) server.RequestHandler {
	return func(handler server.RequestHandler) server.RequestHandler {
		return func(c *server.RequestContext) {
			// Get auth token
			token, err := c.GetBearerToken()
			if err != nil {
				c.WriteError(ErrAuthInvalidToken)
				return
			}

			// Encode body json
			body, err := json.Marshal(
				tokenValidateRequest{
					AccessToken: token,
				},
			)
			if err != nil {
				c.WriteError(errors.ErrServiceUnavailable)
				return
			}

			// Send service request
			res, err := m.httpClientManager.Request(
				m.authServiceEndpoint+"/auth/token/validate",
				client.WithRequestMethod(http.MethodPost),
				client.WithRequestBody(body),
				client.WithRequestContext(c.GetContext()),
			)
			if err != nil {
				c.WriteError(errors.ErrServiceUnavailable)
				return
			}

			// Check status code
			switch res.StatusCode() {
			case 200:
				// Parse response
				var response tokenValidateResult
				if err := json.Unmarshal(res.Body(), &response); err != nil {
					c.WriteError(errors.ErrServiceUnavailable)
					return
				}

				// Set data to ctx
				c.SetUserValue("device", response.Device)
				c.SetUserValue("user", response.User)
				c.SetUserValue("role", response.Role)
				c.SetUserValue("mfa_value", response.Mfa)
				c.SetUserValue("mfa_validation", true)

				// Apply options
				for _, option := range options {
					if err := option.Apply(c, &response); err != nil {
						c.WriteError(err)
						return
					}
				}

				// Check two factor
				mfav, err := c.UserValueBool("mfa_validation")
				if err != nil {
					c.WriteError(err)
					return
				}
				if mfav && response.Mfa {
					c.WriteError(ErrAuth2faRequired)
					return
				}

				handler(c)
			case 400:
				c.WriteError(ErrAuthInvalidToken)
			default:
				c.WriteError(errors.ErrServiceUnavailable)
			}
		}
	}
}

type AuthOption interface {
	Apply(*server.RequestContext, *tokenValidateResult) error
}

// Auth options

// Check role permissions
type authRolesOption struct {
	roles []string
}

func (h authRolesOption) Apply(_ *server.RequestContext, r *tokenValidateResult) error {
	if len(h.roles) > 0 && !slices.Contains(h.roles, r.Role) {
		return ErrAuthInsufficientPermissions
	}
	return nil
}

func WithAuthRolesOption(roles ...string) authRolesOption {
	return authRolesOption{roles}
}

// Without MFA checking
type authMfaOption struct {
}

func (h authMfaOption) Apply(c *server.RequestContext, _ *tokenValidateResult) error {
	c.SetUserValue("mfa_validation", false)
	return nil
}

func WithoutAuthMfaOption() authMfaOption {
	return authMfaOption{}
}

var (
	ErrAuthInsufficientPermissions = errors.New(errors.ErrForbidden, "insufficient_role_permissions")
	ErrAuth2faRequired             = errors.New(errors.ErrUnauthorized, "2fa_required")
	ErrAuthInvalidToken            = errors.New(errors.ErrUnauthorized, "invalid_token")
)
