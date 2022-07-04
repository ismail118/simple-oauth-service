package request

type UserRoleUpdateRequest struct {
	Id   int64  `validate:"required" json:"id"`
	Role string `validate:"required" json:"role"`
}
