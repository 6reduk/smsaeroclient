package apiRequest

import "fmt"

var (
	ErrMethodNotAllowed          = fmt.Errorf("method not allowed")
	ErrPathIsRequired            = fmt.Errorf("path is required")
	ErrMethodIsRequired          = fmt.Errorf("method is required")
	ErrUnableTransformFormParam  = fmt.Errorf("unable transform form parameter")
	ErrUnableTransformQueryParam = fmt.Errorf("unable transform query parameter")
	ErrResponseDataFieldIsEmpty  = fmt.Errorf("response data field is empty")
)
