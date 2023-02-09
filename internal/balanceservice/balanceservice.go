package balanceservice

import (
	"context"

	"github.com/tgnike/yandex-praktikum-diploma/internal/balanceservice/balancerepository"
	"github.com/tgnike/yandex-praktikum-diploma/internal/models"
)

type BalanceService struct {
	repository *balancerepository.BalanceRepository
	accuals    Accruals
}

type Accruals interface {
	GetOrderBalance(order models.OrderNumber) (models.OrderInformation, error)
}

func New(repository *balancerepository.BalanceRepository) *BalanceService {
	return &BalanceService{repository: repository}
}

func (bw *BalanceService) GetBalance(ctx context.Context) float32 {
	return 0
}
func (bw *BalanceService) WithdrawRequest(ctx context.Context) error {
	return nil
}
func (bw *BalanceService) Withdrawals(ctx context.Context) error {
	return nil

}
