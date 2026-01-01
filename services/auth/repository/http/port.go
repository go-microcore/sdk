package adapter

import (
	"context"
)

type Interface interface {
	Auth(ctx context.Context, data AuthData) (*AuthResult, error)
	Auth2fa(ctx context.Context, data Auth2faData) (*Auth2faResult, error)
	TokenRenew(ctx context.Context, data TokenRenewData) (*TokenRenewResult, error)
	TokenValidate(ctx context.Context, data TokenValidateData) (*TokenValidateResult, error)
	UserLogout(ctx context.Context, authToken string) error
	UserLogoutAll(ctx context.Context, authToken string) error
	UserLogoutDevice(ctx context.Context, authToken string, data LogoutDeviceData) error
	UserDevices(ctx context.Context, authToken string) ([]DeviceResult, error)
}
