package domain

import (
	"database/sql"
)

type DataScopeModel struct {
	Id            sql.NullInt64
	UserId        sql.NullInt64
	PrincipalId   sql.NullInt64
	DistributorId sql.NullInt64
	BuyerId       sql.NullInt64
	IsDelete      sql.NullBool
	CreatedAt     sql.NullTime
	UpdatedAt     sql.NullTime
	CreatedBy     sql.NullString
	UpdatedBy     sql.NullString
}
