package request

type ValidateRequest struct {
	OTP string `validate:"required" json:"otp"`
}
