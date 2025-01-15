package domain

import "fmt"

type GenericError struct {
	Code    int
	Message string
	Err     error
}

func NewGenericError(code int, message string, err error) *GenericError {
	return &GenericError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

func (ge *GenericError) Error() string {
	if ge.Err != nil {
		return fmt.Sprintf("%d: %s (%v)", ge.Code, ge.Message, ge.Err)
	}
	return fmt.Sprintf("%d: %s", ge.Code, ge.Message)
}
