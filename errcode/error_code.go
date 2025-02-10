package errcode

import (
	"fmt"

	"github.com/pkg/errors"
)

const (
	// UnknownErrorCode means unknown error
	UnknownErrorCode = 10001
)

type CustomError struct {
	msg  string
	code uint32
}

func NewCustomError(code uint32, msg string) *CustomError {
	return &CustomError{
		msg:  msg,
		code: code,
	}
}

func (e *CustomError) Error() string {
	return e.msg
}

func (e *CustomError) Wrap(errMsg string) *CustomError {
	return &CustomError{
		msg:  fmt.Sprintf("%s: %s", e.msg, errMsg),
		code: e.code,
	}
}

func DecodeError(customErr error) uint32 {
	rootErr := errors.Cause(customErr)
	var ce *CustomError
	if errors.As(rootErr, &ce) {
		return ce.code
	}

	return UnknownErrorCode
}
