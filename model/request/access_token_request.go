package request

type AccessTokenRequest struct {
	AuthorizationCode string `validate:"required" json:"authorization_code"`
	ClientId          int64  `validate:"required" json:"client_id"`
	ClientSecret      string `validate:"required" json:"client_secret"`
}
