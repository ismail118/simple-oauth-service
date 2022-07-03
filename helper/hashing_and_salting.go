package helper

import "golang.org/x/crypto/bcrypt"

func HashAndSalt(pwd string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	PanicIfError(err)

	return string(hashedPassword)
}

func CompareHasPassword(hashPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
}
