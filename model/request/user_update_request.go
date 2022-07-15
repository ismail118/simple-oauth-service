package request

type UserUpdateRequest struct {
	Id            int64  `validate:"required" json:"id"`
	Email         string `validate:"required" json:"email"`
	FirstName     string `validate:"required" json:"first_name"`
	LastName      string `validate:"required" json:"last_name"`
	UserRoleId    int64  `validate:"required" json:"user_role_id"`
	CompanyId     int64  `validate:"required" json:"company_id"`
	PrincipalId   int64  `json:"principal_id"`
	DistributorId int64  `json:"distributor_id"`
	BuyerId       int64  `json:"buyer_id"`
	IsDelete      bool   `json:"is_delete"`
}
