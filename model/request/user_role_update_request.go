package request

type UserRoleUpdateRequest struct {
	Id   int64  `json:"id"`
	Role string `json:"role"`
}
