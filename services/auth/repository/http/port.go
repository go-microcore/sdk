package adapter

import (
	"context"
)

type Interface interface {
	// Devices
	GetDevices(ctx context.Context, authToken string) ([]DeviceResult, error)
	// Logout
	Logout(ctx context.Context, authToken string) error
	LogoutAll(ctx context.Context, authToken string) error
	LogoutDevice(ctx context.Context, authToken string, data LogoutDeviceData) error
	// Roles
	CreateRole(ctx context.Context, authToken string, data CreateRoleData) (*CreateRoleResult, error)
	FilterRoles(ctx context.Context, authToken string, data FilterRolesData) ([]FilterRolesResult, error)
	UpdateRole(ctx context.Context, authToken string, id string, data UpdateRoleData) error
	DeleteRole(ctx context.Context, authToken string, id string) error
	// Rules (HTTP)
	CreateHttpRule(ctx context.Context, authToken string, data CreateHttpRuleData) (*CreateHttpRuleResult, error)
	FilterHttpRules(ctx context.Context, authToken string, data FilterHttpRulesData) ([]FilterHttpRulesResult, error)
	UpdateHttpRule(ctx context.Context, authToken string, id uint, data UpdateHttpRuleData) error
	DeleteHttpRule(ctx context.Context, authToken string, id uint) error
	// Tokens
	Auth(ctx context.Context, data AuthData) (*AuthResult, error)
	Auth2fa(ctx context.Context, data Auth2faData) (*Auth2faResult, error)
	TokenRenew(ctx context.Context, data TokenRenewData) (*TokenRenewResult, error)
	TokenValidate(ctx context.Context, authToken string) (*TokenValidateResult, error)
	TokenAuthorizeHttp(ctx context.Context, authToken string, data TokenAuthorizeHttpData) (*TokenAuthorizeHttpResult, error)
	// Static access tokens
	CreateStaticAccessToken(ctx context.Context, authToken string, data CreateStaticAccessTokenData) (*CreateStaticAccessTokenResult, error)
	FilterStaticAccessTokens(ctx context.Context, authToken string, data FilterStaticAccessTokenData) ([]FilterStaticAccessTokenResult, error)
	DeleteStaticAccessToken(ctx context.Context, authToken string, id string) error
}
