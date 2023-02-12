package authservice

import (
	"context"
	"errors"
	"sync"

	"github.com/tgnike/yandex-praktikum-diploma/internal/models"
)

type UserService struct {
	repository AuthRepository
	auth       TokenRepository
}

type TokenRepository interface {
	Add(userID models.UserID) models.Token
	Get(token models.Token) models.UserID
}

func New(repository AuthRepository) *UserService {
	return &UserService{repository: repository, auth: &SimpleTokenRepository{mu: &sync.RWMutex{}, repo: make(map[models.Token]models.UserID)}}
}

type AuthRepository interface {
	CommitUser(ctx context.Context, username string, password string) (models.UserID, error)
	GetUserID(ctx context.Context, username string, password string) (models.UserID, error)
}

func (u *UserService) Register(ctx context.Context, userJSON *models.UserJSON) (models.Token, error) {
	userID, err := u.repository.CommitUser(ctx, userJSON.Login, userJSON.Password)
	if err != nil {
		return "", err
	}

	token := u.auth.Add(userID)

	return token, nil

}
func (u *UserService) Login(ctx context.Context, userJSON *models.UserJSON) (models.Token, error) {

	userID, err := u.repository.GetUserID(ctx, userJSON.Login, userJSON.Password)

	if err != nil {
		return "", err
	}

	return u.auth.Add(userID), nil
}

func (u *UserService) CheckAuthToken(token string) (models.UserID, error) {

	id := u.auth.Get(models.Token(token))

	if id == "" {
		return models.UserID(""), errors.New("noauth")
	}
	return models.UserID(id), nil

}
