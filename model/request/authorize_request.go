package request

type AuthorizeRequest struct {
	ClientId    int64  `json:"client_id"`
	RedirectUrl string `json:"redirect_url"`
}
