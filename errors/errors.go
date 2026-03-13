package errors

import (
	"encoding/json"
	"fmt"

	_ "go.microcore.dev/framework"
)

type (
	Error struct {
		Message string
		Code    string
	}
)

func (e Error) Error() string {
	return e.Message
}

func (e Error) GetCode() string {
	return e.Code
}

func NewError(message string, code string) error {
	return Error{
		Message: message,
		Code:    code,
	}
}

func UnmarshalError(b []byte) error {
	var e Error
	if err := json.Unmarshal(b, &e); err != nil {
		return fmt.Errorf("unexpected response: %s", b)
	}
	return e
}
