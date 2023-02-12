package balancerepository

import (
	"context"

	"github.com/tgnike/yandex-praktikum-diploma/internal/models"
)

type BalanceRepository struct {
	storage Storage
}

type Storage interface {
	Withdraw(ctx context.Context, order *models.OrderNumber, sum float32, user *models.UserID) error
	GetBalance(ctx context.Context, user *models.UserID) (*models.Balance, error)
	SetBalance()
}

func New(storage Storage) *BalanceRepository {
	return &BalanceRepository{storage: storage}
}

func (r *BalanceRepository) GetBalance(ctx context.Context, user *models.UserID) (*models.Balance, error) {
	return r.storage.GetBalance(ctx, user)
}

func (r *BalanceRepository) Withdraw(ctx context.Context, order *models.OrderNumber, sum float32, user *models.UserID) error {
	return r.storage.Withdraw(ctx, order, sum, user)
}
