package request

type AccessTokenRequest struct {
	Code         string `validate:"required" json:"code"`
	ClientId     int64  `validate:"required" json:"client_id"`
	ClientSecret string `validate:"required" json:"client_secret"`
}
