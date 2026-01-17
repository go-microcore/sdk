package adapter

import (
	"context"
)

type Interface interface {
	Signin(ctx context.Context, data SigninData) (*SigninResult, error)
	Signup(ctx context.Context, data SignupData) (*SignupResult, error)
	Profile(ctx context.Context, authToken string) (*ProfileResult, error)
	TwoFAValidate(ctx context.Context, authToken string, data TwoFAValidateData) (*TwoFAValidateResult, error)
	TwoFASettings(ctx context.Context, authToken string, data TwoFASettingsData) (*TwoFASettingsResult, error)
	TwoFAEnable(ctx context.Context, authToken string, data TwoFAEnableData) error
	TwoFADisable(ctx context.Context, authToken string, data TwoFADisableData) error
	CreateUser(ctx context.Context, authToken string, data CreateUserData) (*CreateUserResult, error)
	FilterUsers(ctx context.Context, authToken string, data FilterUsersData) ([]FilterUsersResult, error)
	UpdateUser(ctx context.Context, authToken string, id string, data UpdateUserData) error
	DeleteUser(ctx context.Context, authToken string, id uint) error
	CreateRole(ctx context.Context, authToken string, data CreateRoleData) (*CreateRoleResult, error)
	FilterRoles(ctx context.Context, authToken string, data FilterRolesData) ([]FilterRolesResult, error)
	UpdateRole(ctx context.Context, authToken string, id string, data UpdateRoleData) error
	DeleteRole(ctx context.Context, authToken string, id string) error
}
