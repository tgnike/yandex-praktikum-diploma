package balancerepository

type BalanceRepository struct {
	storage Storage
}

type Storage interface {
	GetBalance()
	SetBalance()
}

func New(storage Storage) *BalanceRepository {
	return &BalanceRepository{storage: storage}
}
