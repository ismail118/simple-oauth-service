package request

type AuthorizeRequest struct {
	ClientId    int64  `validate:"required" json:"client_id"`
	RedirectUrl string `validate:"required" json:"redirect_url"`
}
