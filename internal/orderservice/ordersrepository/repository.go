package ordersrepository

type OrdersRepository struct {
	storage Storage
}

type Storage interface {
	GetBalance()
	SetBalance()
}

func New(storage Storage) *OrdersRepository {
	return &OrdersRepository{storage: storage}
}
