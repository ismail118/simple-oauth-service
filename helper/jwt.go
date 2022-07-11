package helper

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"simple-oauth-service/constanta"
	"simple-oauth-service/ctx"
	errors2 "simple-oauth-service/errors"
	"time"
)

type MyCustomClaims struct {
	jwt.RegisteredClaims
	Context ctx.Context
}

// NewMyCustomClaims
// Params expiredAt example code time.Duration(1) * time.Hour
func NewMyCustomClaims(Context ctx.Context, issuer string, subject string, expiredAt time.Duration) *MyCustomClaims {
	return &MyCustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    issuer,
			Subject:   subject,
			Audience:  nil,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiredAt)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        "",
		},
		Context: Context,
	}
}

func GenerateJwtToken(claims *MyCustomClaims, signingKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(signingKey))
	if err != nil {
		return "", err
	}

	return ss, nil
}

func ParseJwtTokenToClaims(tokenStr string, signingKey string) (*MyCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(signingKey), nil
	})

	if claim, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return claim, nil
	} else if errors.Is(err, jwt.ErrTokenMalformed) {
		return nil, errors2.NewUnauthorizedError(constanta.InvalidToken)
	} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
		return nil, errors2.NewUnauthorizedError(constanta.ExpiredToken)
	} else {
		return nil, errors2.NewUnauthorizedError(constanta.StatusUnauthorized)
	}
}
