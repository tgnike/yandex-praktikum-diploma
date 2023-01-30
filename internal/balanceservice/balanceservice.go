package balanceservice

import "github.com/tgnike/yandex-praktikum-diploma/internal/balanceservice/balancerepository"

type BalanceService struct {
	repository *balancerepository.BalanceRepository
}

func New(repository *balancerepository.BalanceRepository) *BalanceService {
	return &BalanceService{repository: repository}
}

func (bw *BalanceService) GetBalance() float32 {
	return 0
}
func (bw *BalanceService) WithdrawRequest() error {
	return nil
}
func (bw *BalanceService) Withdrawals() error {
	return nil

}
