package domain

import "time"

type DataScopeModel struct {
	Id            int64
	UserId        int64
	PrincipalId   int64
	DistributorId int64
	BuyerId       int64
	IsDelete      bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
	CreatedBy     string
	UpdatedBy     string
}
