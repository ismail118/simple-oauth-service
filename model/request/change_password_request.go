package request

type ChangePasswordRequest struct {
	Id          int64  `validate:"required" json:"id"`
	OldPassword string `validate:"required" json:"old_password"`
	NewPassword string `validate:"required" json:"new_password"`
}
