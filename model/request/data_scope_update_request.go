package request

type DataScopeUpdateRequest struct {
	Id            int64 `json:"id"`
	UserId        int64 `json:"user_id"`
	PrincipalId   int64 `json:"principal_id"`
	DistributorId int64 `json:"distributor_id"`
	BuyerId       int64 `json:"buyer_id"`
	IsDelete      bool  `json:"is_delete"`
}
