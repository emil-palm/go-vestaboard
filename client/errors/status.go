package errors

import (
	"fmt"
)

type StatusError struct {
	StatusCode int
	Err        error
}

func (re StatusError) Error() string {
	return fmt.Sprintf("%d: %s", re.StatusCode, re.Err)
}

func NewStatusError(statusCode int, err error) StatusError {
	return StatusError{
		StatusCode: statusCode,
		Err:        err,
	}
}

func StatusErrorf(statusCode int, format string, args ...interface{}) StatusError {
	return NewStatusError(statusCode, fmt.Errorf(format, args...))
}
