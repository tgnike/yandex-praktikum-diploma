package ordersrepository

import (
	"context"

	"github.com/tgnike/yandex-praktikum-diploma/internal/models"
)

type OrdersRepository struct {
	storage Storage
}

type Storage interface {
	GetOrderByNumber(ctx context.Context, order string, userID string) error
	CommitOrderNumber(ctx context.Context, order string, status string, balance float32, userID string) error
	GetUserOrders(ctx context.Context, user string) ([]models.OrderInformation, error)
	UpdateOrder(ctx context.Context, order string, status string, balance float32) error
}

func New(storage Storage) *OrdersRepository {
	return &OrdersRepository{storage: storage}
}

func (r *OrdersRepository) GetOrderByNumber(ctx context.Context, order models.OrderNumber, user models.UserID) error {

	return r.storage.GetOrderByNumber(ctx, string(order), string(user))

}
func (r *OrdersRepository) CommitOrderNumber(ctx context.Context, order models.OrderNumber, status models.OrderStatus, user models.UserID) error {

	return r.storage.CommitOrderNumber(ctx, string(order), string(status), 0, string(user))

}

func (r *OrdersRepository) GetUserOrders(ctx context.Context, user models.UserID) ([]models.OrderInformation, error) {
	return r.storage.GetUserOrders(ctx, string(user))
}

func (r *OrdersRepository) UpdateOrderAccruals(ctx context.Context, order *models.OrderInformation) error {
	return r.storage.UpdateOrder(ctx, string(order.Order), string(order.Status), order.Balance)
}
