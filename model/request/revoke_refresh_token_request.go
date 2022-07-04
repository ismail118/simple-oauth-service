package request

type RevokeRefreshTokenRequest struct {
	UserId int64 `validate:"required" json:"user_id"`
}
