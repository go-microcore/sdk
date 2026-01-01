package adapter

import "go.microcore.dev/framework/errors"

var (
	ErrInvalidToken     = errors.New(errors.ErrBadRequest, "invalid_token")
	ErrTokenAlreadyUsed = errors.New(errors.ErrBadRequest, "token_already_used")
)
