package errcode

import (
	"errors"
	"fmt"
)

var (
	ErrResultsNotFound = errors.New("results not found")
)

type Error struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

func New(code int64, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

func (e Error) Error() string {
	return fmt.Sprintf("code: %d, message: %s", e.Code, e.Message)
}

func FromError(err error) *Error {
	if err == nil {
		return nil
	}
	if e := new(Error); errors.As(err, &e) {
		return e
	}
	return nil
}

func Is(err error, code int64) bool {
	if e := FromError(err); e != nil {
		return e.Code == code
	}
	return false
}
