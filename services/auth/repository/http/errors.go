package adapter

import "go.microcore.dev/framework/errors"

var (
	ErrInvalidDevice           = errors.New(errors.ErrBadRequest, "invalid_device")
	ErrInvalidRoleId           = errors.New(errors.ErrBadRequest, "invalid_role_id")
	ErrInvalidRoleName         = errors.New(errors.ErrBadRequest, "invalid_role_name")
	ErrInvalidRoleDescription  = errors.New(errors.ErrBadRequest, "invalid_role_description")
	ErrRoleExistId             = errors.New(errors.ErrBadRequest, "role_exist_id")
	ErrRoleNotFound            = errors.New(errors.ErrBadRequest, "role_not_found")
	ErrInvalidPath             = errors.New(errors.ErrBadRequest, "invalid_path")
	ErrInvalidMethods          = errors.New(errors.ErrBadRequest, "invalid_methods")
	ErrInvalidMfa              = errors.New(errors.ErrBadRequest, "invalid_mfa")
	ErrRuleNotFound            = errors.New(errors.ErrBadRequest, "rule_not_found")
	ErrRuleExist               = errors.New(errors.ErrBadRequest, "rule_exist")
	ErrInvalidToken            = errors.New(errors.ErrBadRequest, "invalid_token")
	ErrTokenAlreadyUsed        = errors.New(errors.ErrBadRequest, "token_already_used")
	ErrInvalidRoles            = errors.New(errors.ErrBadRequest, "invalid_roles")
	ErrInvalidDescription      = errors.New(errors.ErrBadRequest, "invalid_description")
	ErrStaticTokenNotFound     = errors.New(errors.ErrBadRequest, "static_token_not_found")
	ErrStaticTokenExist        = errors.New(errors.ErrBadRequest, "static_token_exist")
	ErrInsufficientPermissions = errors.New(errors.ErrForbidden, "insufficient_permissions")
	ErrInvalidId               = errors.New(errors.ErrForbidden, "invalid_id")
)
