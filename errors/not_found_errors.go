package errors

type NotFoundError struct {
	MessageError string
}

func NewNotFoundError(error string) NotFoundError {
	return NotFoundError{MessageError: error}
}

func (err NotFoundError) Error() string {
	return err.MessageError
}
