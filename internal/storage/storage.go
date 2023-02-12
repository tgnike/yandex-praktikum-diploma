package storage

import (
	"context"

	"github.com/tgnike/yandex-praktikum-diploma/internal/models"
)

type Storage struct {
	db DB
}

type DB interface {
	Init()
	// Close()
	// Ping() error
	GetBalance(ctx context.Context, user string) (*models.Balance, error)
	// PostBalance(user string, balance float32)
	// PostOrder(order string) error
	// GetOrdersInformation() ([]string, error)
	GetOrder(ctx context.Context, order string, userID string) error
	GetUserOrders(ctx context.Context, userID string, orders models.OrderContainerInterface) error
	CommitOrder(ctx context.Context, order string, status string, balance float32, userID string) error
	UpdateOrder(ctx context.Context, order string, status string, balance float32) error

	GetUser(ctx context.Context, username string, password string) (string, error)
	StoreUser(ctx context.Context, username string, password string) (string, error)

	Withdraw(ctx context.Context, order string, sum float32, user string) error
}

const dbTimeout = 5

func NewStore(db DB) *Storage {
	return &Storage{db: db}
}

func (s *Storage) CommitUser(ctx context.Context, username string, password string) (string, error) {
	return s.db.StoreUser(ctx, username, password)
}

func (s *Storage) GetUser(ctx context.Context, username string, password string) (string, error) {
	return s.db.GetUser(ctx, username, password)
}

func (s *Storage) GetOrderByNumber(ctx context.Context, order string, userID string) error {
	return s.db.GetOrder(ctx, order, userID)
}

func (s *Storage) CommitOrderNumber(ctx context.Context, order string, status string, balance float32, userID string) error {
	return s.db.CommitOrder(ctx, order, status, balance, userID)
}

func (s *Storage) GetUserOrders(ctx context.Context, user string, orders models.OrderContainerInterface) error {
	err := s.db.GetUserOrders(ctx, user, orders)
	if err != nil {
		return err
	}
	return nil

}

func (s *Storage) UpdateOrder(ctx context.Context, order string, status string, balance float32) error {
	return s.db.UpdateOrder(ctx, order, status, balance)
}

func (s *Storage) Withdraw(ctx context.Context, order *models.OrderNumber, sum float32, user *models.UserID) error {
	return s.db.Withdraw(ctx, string(*order), sum, string(*user))
}

func (s *Storage) GetBalance(ctx context.Context, user *models.UserID) (*models.Balance, error) {

	return s.db.GetBalance(ctx, string(*user))

}

func (s *Storage) SetBalance() {}
