package balanceservice

import (
	"context"

	"github.com/tgnike/yandex-praktikum-diploma/internal/models"
	"github.com/tgnike/yandex-praktikum-diploma/internal/server"
)

type BalanceService struct {
	repository BalanceRepository
	accuals    Accruals
}

type Accruals interface {
	GetOrderBalance(order models.OrderNumber) (models.OrderInformation, error)
}

type BalanceRepository interface {
	Withdraw(ctx context.Context, order *models.OrderNumber, sum float32, user *models.UserID) error
}

func New(repository BalanceRepository) *BalanceService {
	return &BalanceService{repository: repository}
}

func (bw *BalanceService) GetBalance(ctx context.Context) float32 {
	return 0
}
func (bw *BalanceService) WithdrawRequest(ctx context.Context, withdrawal *models.WithdrawalRequest, user *models.UserID) error {

	orderNumber := &withdrawal.Order
	err := orderNumber.Check()

	if err != nil {
		return server.ErrOrderNumberFormat
	}

	return bw.repository.Withdraw(ctx, &withdrawal.Order, withdrawal.Sum, user)
}
func (bw *BalanceService) Withdrawals(ctx context.Context) error {
	return nil

}
