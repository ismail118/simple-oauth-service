package domain

import (
	"database/sql"
)

type UserRoleModel struct {
	Id        sql.NullInt64
	Role      sql.NullString
	CreatedAt sql.NullTime
}
