package helper

import (
	"math/rand"
)

const letterNumberBytes = "1234567890"
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}

func RandRandomStringNumber(max int) string {
	b := make([]byte, max)
	for i := range b {
		b[i] = letterNumberBytes[rand.Int63()%int64(len(letterNumberBytes))]
	}
	return string(b)
}
