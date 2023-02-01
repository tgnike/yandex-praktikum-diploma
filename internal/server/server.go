package server

import (
	"context"

	"github.com/tgnike/yandex-praktikum-diploma/internal/models"
)

type Server struct {
	Users   Users
	Orders  Orders
	Balance Balance
}

type UserContextKey string

const UserContext UserContextKey = "usrctxkey"

type Users interface {
	Register(userJSON *models.UserJSON) (models.Token, error)
	Login(userJSON *models.UserJSON) (models.Token, error)
	CheckAuthToken(token string) (models.UserID, error)
}

type Orders interface {
	PostOrder(ctx context.Context, order models.OrderNumber, user models.UserID) error
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
