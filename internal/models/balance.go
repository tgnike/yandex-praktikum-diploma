package models

import "time"

type WithdrawalRequest struct {
	Order OrderNumber `json:"order"`
	Sum   float32     `json:"sum"`
}

type Withdrawal struct {
	Order OrderNumber `json:"order"`
	Sum   float32     `json:"sum"`
	Date  time.Time   `json:"processed_at"`
}

type Balance struct {
	Current   float32 `json:"current"`
	Withdrawn float32 `json:"withdrawn"`
}
