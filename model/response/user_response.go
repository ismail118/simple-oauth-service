package response

import "time"

type UserResponse struct {
	Id            int64     `json:"id"`
	Email         string    `json:"email"`
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	UserRoleId    int64     `json:"user_role_id"`
	CompanyId     int64     `json:"company_id"`
	PrincipalId   int64     `json:"principal_id"`
	DistributorId int64     `json:"distributor_id"`
	BuyerId       int64     `json:"buyer_id"`
	TokenVersion  int64     `json:"token_version"`
	IsVerified    bool      `json:"is_verified"`
	IsDelete      bool      `json:"is_delete"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	CreatedBy     string    `json:"created_by"`
	UpdatedBy     string    `json:"updated_by"`
}
