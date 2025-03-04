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

// IsServiceUnavailable 50001 Service temporarily unavailable, try again
func IsServiceUnavailable(err error) bool {
	return Is(err, 50001)
}

// IsRateLimitReached 50011 Rate limit reached. Please refer to API documentation and throttle requests accordingly
func IsRateLimitReached(err error) bool {
	return Is(err, 50011)
}

// IsParameterCannotBeEmpty 50014 Parameter {param0} cannot be empty
func IsParameterCannotBeEmpty(err error) bool {
	return Is(err, 50014)
}

// IsSystemError 50026 System error. Try again later
func IsSystemError(err error) bool {
	return Is(err, 50026)
}

// IsInvalidSignature 50113 Invalid signature
func IsInvalidSignature(err error) bool {
	return Is(err, 50113)
}

// IsParameterError 51000 Parameter {param0} error
func IsParameterError(err error) bool {
	return Is(err, 51000)
}

// IsRepeatedRequest 80000 Repeated request
func IsRepeatedRequest(err error) bool {
	return Is(err, 80000)
}
