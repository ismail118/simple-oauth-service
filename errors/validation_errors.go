package errors

import "fmt"

type ValidationErrors struct {
	MessageError string
}

func NewValidationErrors(error string) ValidationErrors {
	return ValidationErrors{MessageError: error}
}

func NewValidationEmptyErrors(param string) ValidationErrors {
	return ValidationErrors{MessageError: fmt.Sprintf("Error empty %s", param)}
}

func (err ValidationErrors) Error() string {
	return err.MessageError
}
