package domain

import (
	"database/sql"
)

type UserModel struct {
	Id            sql.NullInt64
	Email         sql.NullString
	Password      sql.NullString
	FirstName     sql.NullString
	LastName      sql.NullString
	UserRoleId    sql.NullInt64
	CompanyId     sql.NullInt64
	PrincipalId   sql.NullInt64
	DistributorId sql.NullInt64
	BuyerId       sql.NullInt64
	TokenVersion  sql.NullInt64
	IsVerified    sql.NullBool
	IsDelete      sql.NullBool
	CreatedAt     sql.NullTime
	UpdatedAt     sql.NullTime
	CreatedBy     sql.NullString
	UpdatedBy     sql.NullString
}
