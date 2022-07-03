package response

import "time"

type DataScopeResponse struct {
	Id            int64     `json:"id"`
	UserId        int64     `json:"user_id"`
	PrincipalId   int64     `json:"principal_id"`
	DistributorId int64     `json:"distributor_id"`
	BuyerId       int64     `json:"buyer_id"`
	IsDelete      bool      `json:"is_delete"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	CreatedBy     string    `json:"created_by"`
	UpdatedBy     string    `json:"updated_by"`
}
