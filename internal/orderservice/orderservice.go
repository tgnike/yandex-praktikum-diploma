package orderservice

import (
	"context"

	"github.com/tgnike/yandex-praktikum-diploma/internal/models"
	"github.com/tgnike/yandex-praktikum-diploma/internal/orderservice/ordersrepository"
)

type OrderService struct {
	repository *ordersrepository.OrdersRepository
}

func New(repository *ordersrepository.OrdersRepository) *OrderService {
	return &OrderService{repository: repository}
}

func (ow *OrderService) PostOrder(ctx context.Context, order models.OrderNumber, user models.UserID) error {

	// Проверка контрольной суммы номера заказа
	err := order.Check()

	if err != nil {
		return err
	}

	err = checkOrderAlreadyLoaded(ctx, order, user)

	if err != nil {
		return err
	}

	err = storeOrder(ctx, order, user)

	if err != nil {
		return err
	}

	return nil

}

func storeOrder(ctx context.Context, order models.OrderNumber, user models.UserID) error {
	return nil
}

func checkOrderAlreadyLoaded(ctx context.Context, order models.OrderNumber, user models.UserID) error {
	return nil
}

func (ow *OrderService) GetOrdersInformation() ([]models.OrderInformation, error) {
	return make([]models.OrderInformation, 0), nil
}
