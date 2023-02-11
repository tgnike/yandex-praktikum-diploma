package orderservice

import (
	"context"
	"log"
	"time"

	"github.com/tgnike/yandex-praktikum-diploma/internal/models"
	"github.com/tgnike/yandex-praktikum-diploma/internal/server"
)

type OrderService struct {
	repository OrderRepository
	accruals   Accruals
}

func New(repository OrderRepository, accruals Accruals) *OrderService {
	return &OrderService{repository: repository, accruals: accruals}
}

type OrderRepository interface {
	GetOrderByNumber(ctx context.Context, order models.OrderNumber, user models.UserID) error
	CommitOrderNumber(ctx context.Context, order models.OrderNumber, status models.OrderStatus, user models.UserID) error
	GetUserOrders(ctx context.Context, user models.UserID) ([]*models.OrderInformation, error)
	UpdateOrderAccruals(ctx context.Context, order *models.AccrualInformation) error
}

type Accruals interface {
	Register(order models.OrderNumber)
	Check(order models.OrderNumber)
	GetUpdates() chan *models.AccrualInformation
}

func (os *OrderService) PostOrder(ctx context.Context, order string, user models.UserID) error {

	// Проверка контрольной суммы номера заказа
	orderNumber := models.OrderNumber(order)
	err := orderNumber.Check()

	if err != nil {
		return server.ErrOrderNumberFormat
	}

	err = os.repository.CommitOrderNumber(ctx, orderNumber, models.NEW, user)

	if err != nil {
		return err
	}

	os.accruals.Check(orderNumber)
	// TODO
	//go os.checkWithTimeout(orderNumber)

	return nil

}

func (os *OrderService) checkWithTimeout(order models.OrderNumber) {

	timer := time.NewTimer(time.Duration(time.Second))
	<-timer.C
	os.accruals.Check(order)

}

func (os *OrderService) GetOrdersInformation(ctx context.Context, user models.UserID) ([]*models.OrderInformation, error) {

	orders, err := os.repository.GetUserOrders(ctx, user)

	if err != nil {
		return make([]*models.OrderInformation, 0), err
	}

	return orders, nil
}

func (os *OrderService) UpdateAccrualInformation(ctx context.Context) {

	ctxUpdate, cancelUpdate := context.WithCancel(ctx)

	go func(ctx context.Context, c chan *models.AccrualInformation) {
		for {
			select {
			case orderInfo := <-c:

				os.repository.UpdateOrderAccruals(ctx, orderInfo)

				log.Print(orderInfo)

			case <-ctx.Done():
				return

			}

		}
	}(ctxUpdate, os.accruals.GetUpdates())

	for {
		select {
		case <-ctx.Done():
			cancelUpdate()
			return
		}

	}

}
