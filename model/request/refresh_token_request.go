package request

type RefreshTokenRequest struct {
	RefreshToken string `validate:"required" json:"refresh_token"`
}
