package request

type ClientCreateRequest struct {
	ApplicationName string `validate:"required" json:"application_name"`
}
