package errors

type ForbiddenError struct {
	MessageError string
}

func NewForbiddenError(error string) ForbiddenError {
	return ForbiddenError{MessageError: error}
}

func (err ForbiddenError) Error() string {
	return err.MessageError
}
