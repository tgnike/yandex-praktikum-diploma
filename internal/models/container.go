package models

import "time"

type OrderContainerInterface interface {
	Add(order string, balance float32, status string, date time.Time)
	Value() []*OrderInformation
}

type OrderContainer struct {
	orders []*OrderInformation
}

func (oc *OrderContainer) Add(order string, balance float32, status string, date time.Time) {

	oc.orders = append(oc.orders, NewOrderInfo(order, balance, status, date))

}

func (oc *OrderContainer) Value() []*OrderInformation {

	return oc.orders

}
