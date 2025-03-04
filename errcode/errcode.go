// Copyright (c) 2024-NOW imzhongqi <imzhongqi@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

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
