package models

type WithdrawalRequest struct {
	Order OrderNumber `json:"order"`
	Sum   float32     `json:"sum"`
}

type Balance struct {
	Current   float32 `json:"current"`
	Withdrawn float32 `json:"withdrawn"`
}
