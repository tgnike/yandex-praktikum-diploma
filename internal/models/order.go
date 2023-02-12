package models

import (
	"strconv"
	"time"

	"github.com/theplant/luhn"
)

type OrderInformation struct {
	Order   OrderNumber `json:"order"`
	Status  OrderStatus `json:"status"`
	Balance float32     `json:"accrual,omitempty"`
	Date    time.Time   `json:"uploaded_at"`
}

func NewOrderInfo(order string, balance float32, status string, date time.Time) *OrderInformation {

	return &OrderInformation{Order: OrderNumber(order), Status: OrderStatus(status), Balance: balance, Date: date}

}

type WithdrawalRequest struct {
	Order OrderNumber `json:"order"`
	Sum   float32     `json:"sum"`
}

type AccrualInformation struct {
	Order   OrderNumber
	Status  OrderStatus
	Accrual float32
	Date    time.Time
}

type OrderStatus string

const (
	NEW        OrderStatus = "NEW"
	INVALID    OrderStatus = "INVALID"
	PROCESSING OrderStatus = "PROCESSING"
	PROCESSED  OrderStatus = "PROCESSED"
	REGISTERED OrderStatus = "REGISTERED"
)

type OrderNumber string

func (o *OrderNumber) Check() error {

	i, err := strconv.Atoi(string(*o))

	if err != nil {
		return &InvalidOrderNumber{}
	}

	good := luhn.Valid(i)

	if good {
		return nil
	}

	return &InvalidOrderNumber{}
}

type InvalidOrderNumber struct {
}

func (e *InvalidOrderNumber) Error() string {

	return "invalid order number"

}
