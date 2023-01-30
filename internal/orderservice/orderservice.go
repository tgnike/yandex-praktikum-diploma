package orderservice

import (
	"github.com/tgnike/yandex-praktikum-diploma/internal/models"
	"github.com/tgnike/yandex-praktikum-diploma/internal/orderservice/ordersrepository"
)

type OrderService struct {
	repository *ordersrepository.OrdersRepository
}

func New(repository *ordersrepository.OrdersRepository) *OrderService {
	return &OrderService{repository: repository}
}

func (ow *OrderService) PostOrder(order models.OrderNumber) error {
	return nil
}

func (ow *OrderService) GetOrdersInformation() ([]models.OrderInformation, error) {
	return make([]models.OrderInformation, 0), nil
}
