package httpApiClient

import "fmt"

var (
	ErrExecutionRequest = fmt.Errorf("execution request error")
	ErrReadingBody      = fmt.Errorf("reading response body error")
)

type ApiError struct {
	Err     error
	Message string
	Code    int
}

func NewApiError(err error, code int, message string) *ApiError {
	return &ApiError{
		Err:     err,
		Code:    code,
		Message: message,
	}
}

func (e *ApiError) Error() string {
	return fmt.Sprintf(
		"error occurred on api-service side, code: %d, message: %s, reason: %v",
		e.Code,
		e.Message,
		e.Err,
	)
}

func (e *ApiError) Unwrap() error {
	return e.Err
}

type HandleResponseError struct {
	Err     error
	Message string
	Code    int
}

func NewHandleResponseError(err error, code int, message string) *HandleResponseError {
	return &HandleResponseError{
		Err:     err,
		Code:    code,
		Message: message,
	}
}

func (hre *HandleResponseError) Error() string {
	return fmt.Sprintf(
		"error occurred during process response: %d, message: %s, reason: <%s>",
		hre.Code,
		hre.Message,
		hre.Err,
	)
}

func (hre *HandleResponseError) Unwrap() error {
	return hre.Err
}
