package request

type UserUpdateRequest struct {
	Id            int64  `json:"id"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	UserRoleId    int64  `json:"user_role_id"`
	CompanyId     int64  `json:"company_id"`
	PrincipalId   int64  `json:"principal_id"`
	DistributorId int64  `json:"distributor_id"`
	BuyerId       int64  `json:"buyer_id"`
	IsVerified    bool   `json:"is_verified"`
	IsDelete      bool   `json:"is_delete"`
}
