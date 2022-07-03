package helper

import (
	"errors"
	"io"
)

func PanicIfError(err error) {
	if err != nil && !errors.Is(err, io.EOF) {
		panic(err)
	}
}
