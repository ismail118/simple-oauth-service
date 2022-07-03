package domain

import "time"

type UserModel struct {
	Id            int64
	Email         string
	Password      string
	FirstName     string
	LastName      string
	UserRoleId    int64
	CompanyId     int64
	PrincipalId   int64
	DistributorId int64
	BuyerId       int64
	TokenVersion  int64
	IsVerified    bool
	IsDelete      bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
	CreatedBy     string
	UpdatedBy     string
}
