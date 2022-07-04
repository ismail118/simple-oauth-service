package request

type ClientUpdateRequest struct {
	Id              int64  `validate:"required" json:"id"`
	ApplicationName string `validate:"required" json:"application_name"`
	IsDelete        bool   `validate:"required" json:"is_delete"`
}
