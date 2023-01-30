package models

type OrderInformation struct {
	Order   OrderNumber
	Status  string
	Balance float32
}

type OrderNumber string
