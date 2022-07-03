package request

type ClientUpdateRequest struct {
	Id              int64  `json:"id"`
	ApplicationName string `json:"application_name"`
	ClientSecret    string `json:"client_secret"`
	IsDelete        bool   `json:"is_delete"`
}
