package balanceservice

import (
	"context"

	"github.com/tgnike/yandex-praktikum-diploma/internal/models"
	"github.com/tgnike/yandex-praktikum-diploma/internal/server"
)

type BalanceService struct {
	repository BalanceRepository
}

type BalanceRepository interface {
	Withdraw(ctx context.Context, order *models.OrderNumber, sum float32, user *models.UserID) error
	GetBalance(ctx context.Context, user *models.UserID) (*models.Balance, error)
	Withdrawals(ctx context.Context, user *models.UserID) ([]*models.Withdrawal, error)
}

func New(repository BalanceRepository) *BalanceService {
	return &BalanceService{repository: repository}
}

func (bw *BalanceService) GetBalance(ctx context.Context, user *models.UserID) (*models.Balance, error) {
	return bw.repository.GetBalance(ctx, user)
}
func (bw *BalanceService) WithdrawRequest(ctx context.Context, withdrawal *models.WithdrawalRequest, user *models.UserID) error {

	orderNumber := &withdrawal.Order
	err := orderNumber.Check()

	if err != nil {
		return server.ErrOrderNumberFormat
	}

	return bw.repository.Withdraw(ctx, &withdrawal.Order, withdrawal.Sum, user)
}
func (bw *BalanceService) Withdrawals(ctx context.Context, user *models.UserID) ([]*models.Withdrawal, error) {
	return bw.repository.Withdrawals(ctx, user)
}
