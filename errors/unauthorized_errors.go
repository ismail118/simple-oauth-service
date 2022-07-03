package errors

type UnauthorizedError struct {
	MessageError string
}

func NewUnauthorizedError(error string) UnauthorizedError {
	return UnauthorizedError{MessageError: error}
}

func (err UnauthorizedError) Error() string {
	return err.MessageError
}
