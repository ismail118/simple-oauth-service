package domain

import (
	"database/sql"
)

type ClientModel struct {
	Id              sql.NullInt64
	UserId          sql.NullInt64
	ApplicationName sql.NullString
	ClientSecret    sql.NullString
	IsDelete        sql.NullBool
	CreatedAt       sql.NullTime
	UpdatedAt       sql.NullTime
	CreatedBy       sql.NullString
	UpdatedBy       sql.NullString
}
