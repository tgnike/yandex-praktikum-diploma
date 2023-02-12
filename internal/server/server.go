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
	Register(ctx context.Context, userJSON *models.UserJSON) (models.Token, error)
	Login(ctx context.Context, userJSON *models.UserJSON) (models.Token, error)
	CheckAuthToken(token string) (models.UserID, error)
}

type Orders interface {
	PostOrder(ctx context.Context, order string, user models.UserID) error
	GetOrdersInformation(ctx context.Context, user models.UserID) ([]*models.OrderInformation, error)
	UpdateAccrualInformation(ctx context.Context)
}

type Balance interface {
	GetBalance(ctx context.Context, user *models.UserID) (*models.Balance, error)
	WithdrawRequest(ctx context.Context, withdrawal *models.WithdrawalRequest, user *models.UserID) error
	Withdrawals(ctx context.Context, user *models.UserID) error
}

func New(users Users, orders Orders, balance Balance) *Server {

	return &Server{Users: users, Orders: orders, Balance: balance}

}
