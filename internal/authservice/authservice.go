package authservice

import (
	"errors"

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
	return &UserService{repository: repository}
}

type AuthRepository interface {
	StoreUser(username string, password string) (models.UserID, error)
	GetUserID(username string, password string) (models.UserID, error)
}

func (u *UserService) Register(userJSON *models.UserJSON) (models.Token, error) {
	userID, err := u.repository.StoreUser(userJSON.Login, userJSON.Password)
	if err != nil {
		return "", err
	}

	token := u.auth.Add(userID)

	return token, nil

}
func (u *UserService) Login(userJSON *models.UserJSON) (models.Token, error) {

	userID, err := u.repository.GetUserID(userJSON.Login, userJSON.Password)

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
