package authservice

import authrepository "github.com/tgnike/yandex-praktikum-diploma/internal/authservice/authrepository"

type UserService struct {
	repository *authrepository.AuthRepository
}

func New(repository *authrepository.AuthRepository) *UserService {
	return &UserService{repository: repository}
}

func (u *UserService) Register(username string, password string) error {
	return nil
}
func (u *UserService) Login(username string, password string) (string, error) {
	return "", nil
}
