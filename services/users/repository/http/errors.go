package adapter

import "go.microcore.dev/framework/errors"

var (
	ErrInvalidCredentials = errors.New(errors.ErrUnauthorized, "invalid_credentials")
	ErrInvalidLogin       = errors.New(errors.ErrBadRequest, "invalid_login")
	ErrInvalidPassword    = errors.New(errors.ErrBadRequest, "invalid_password")
	ErrInvalidUsername    = errors.New(errors.ErrBadRequest, "invalid_username")
	ErrInvalidEmail       = errors.New(errors.ErrBadRequest, "invalid_email")
	ErrInvalidName        = errors.New(errors.ErrBadRequest, "invalid_name")
	ErrExistEmail         = errors.New(errors.ErrBadRequest, "user_exist_email")
	ErrExistUsername      = errors.New(errors.ErrBadRequest, "user_exist_username")
	ErrMfaDisabled        = errors.New(errors.ErrBadRequest, "mfa_disabled")
	ErrMfaEnabled         = errors.New(errors.ErrBadRequest, "mfa_enabled")
	ErrInvalidToken       = errors.New(errors.ErrBadRequest, "invalid_token")
	ErrRoleNotFound       = errors.New(errors.ErrBadRequest, "role_not_found")
	ErrNotFound           = errors.New(errors.ErrBadRequest, "user_not_found")
	ErrUserIsUsed         = errors.New(errors.ErrBadRequest, "user_is_used")
	ErrInvalidRoles       = errors.New(errors.ErrBadRequest, "invalid_roles")
)
