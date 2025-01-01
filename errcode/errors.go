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
