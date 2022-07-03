package helper

import (
	"simple-oauth-service/constanta"
	"simple-oauth-service/ctx"
	"simple-oauth-service/errors"
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

	return errors.ForbiddenError{
		MessageError: constanta.Status403,
	}
}
