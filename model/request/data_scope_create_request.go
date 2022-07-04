package request

type DataScopeCreateRequest struct {
	UserId        int64 `validate:"required" json:"user_id"`
	PrincipalId   int64 `json:"principal_id"`
	DistributorId int64 `json:"distributor_id"`
	BuyerId       int64 `json:"buyer_id"`
}
