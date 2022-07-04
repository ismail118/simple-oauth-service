package request

type UserRoleCreateRequest struct {
	Role string `validate:"required" json:"role"`
}
