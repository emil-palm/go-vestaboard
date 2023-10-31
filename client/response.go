package client

import "fmt"

type APIError struct {
	StatusCode int
	Err        error
}

func (ae *APIError) Error() string {
	return fmt.Sprintf("%d: %s", ae.StatusCode, ae.Err)
}

func WrapAPIError(err error, statusCode int) *APIError {
	return &APIError{
		StatusCode: statusCode,
		Err:        err,
	}
}
