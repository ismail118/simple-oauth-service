package helper

import (
	"net/http"
	"simple-oauth-service/constanta"
	"simple-oauth-service/ctx"
	errors2 "simple-oauth-service/errors"
	"strings"
)

func CheckRoles(ctx ctx.Context, roles ...string) error {
	if len(roles) == 0 {
		return nil
	}

	rolesStr := strings.Join(roles, " ")

	if strings.Contains(rolesStr, ctx.UserRole.Role) {
		return nil
	}

	return errors2.ForbiddenError{
		MessageError: constanta.Status403,
	}
}

func ValidateRefreshToken(r *http.Request) (*http.Cookie, bool) {
	c, err := r.Cookie("jid")
	if err != nil {
		return nil, false
	}

	_, err = ParseJwtTokenToClaims(c.Value, constanta.SecretKey)
	if err != nil {
		return nil, false
	}

	return c, true
}
