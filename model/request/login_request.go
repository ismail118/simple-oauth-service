package request

type LoginRequest struct {
	Email    string `validate:"email" json:"email"`
	Password string `validate:"required" json:"password"`
}
