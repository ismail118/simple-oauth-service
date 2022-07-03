package response

import "time"

type ClientResponse struct {
	Id              int64     `json:"id"`
	UserId          int64     `json:"user_id"`
	ApplicationName string    `json:"application_name"`
	ClientSecret    string    `json:"client_secret"`
	IsDelete        bool      `json:"is_delete"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	CreatedBy       string    `json:"created_by"`
	UpdatedBy       string    `json:"updated_by"`
}
