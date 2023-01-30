package authservice

import (
	authrepository "github.com/tgnike/yandex-praktikum-diploma/internal/authservice/authrepository"
	"github.com/tgnike/yandex-praktikum-diploma/internal/models"
)

type UserService struct {
	repository *authrepository.AuthRepository
}

func New(repository *authrepository.AuthRepository) *UserService {
	return &UserService{repository: repository}
}

func (u *UserService) Register(userJSON *models.UserJSON) error {
	return nil
}
func (u *UserService) Login(userJSON *models.UserJSON) (string, error) {
	return "", nil
}
