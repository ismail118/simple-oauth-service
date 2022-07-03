package request

type AccessTokenRequest struct {
	AuthorizationCode string `json:"authorization_code"`
	ClientId          int64  `json:"client_id"`
	ClientSecret      string `json:"client_secret"`
}
