package models

import (
	"strconv"

	"github.com/theplant/luhn"
)

type OrderInformation struct {
	Order   OrderNumber
	Status  OrderStatus
	Balance float32
}

type OrderStatus string

const (
	NEW        OrderStatus = "NEW"
	INVALID    OrderStatus = "INVALID"
	PROCESSING OrderStatus = "PROCESSING"
	PROCESSED  OrderStatus = "PROCESSED"
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
