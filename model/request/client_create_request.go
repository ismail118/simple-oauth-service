package request

type ClientCreateRequest struct {
	ApplicationName string `json:"application_name"`
	ClientSecret    string `json:"client_secret"`
	IsDelete        bool   `json:"is_delete"`
}
