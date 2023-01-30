package server

import (
	"github.com/tgnike/yandex-praktikum-diploma/internal/models"
)

type Server struct {
	Users   Users
	Orders  Orders
	Balance Balance
}

type Users interface {
	Register(username string, password string) error
	Login(username string, password string) (string, error)
}

type Orders interface {
	PostOrder(order models.OrderNumber) error
	GetOrdersInformation() ([]models.OrderInformation, error)
}

type Balance interface {
	GetBalance() float32
	WithdrawRequest() error
	Withdrawals() error
}

func New(users Users, orders Orders, balance Balance) *Server {

	return &Server{Users: users, Orders: orders, Balance: balance}

}
