package authrepository

import "github.com/tgnike/yandex-praktikum-diploma/internal/models"

type AuthRepository struct {
	Storage Storage
}

type Storage interface {
	CommitUser(username string, password string) (string, error)
	GetUser(username string, password string) (string, error)
}

func New(storage Storage) *AuthRepository {
	return &AuthRepository{Storage: storage}
}

func (r *AuthRepository) StoreUser(username string, password string) (models.UserID, error) {

	id, err := r.Storage.CommitUser(username, password)
	if err != nil {
		return "", err
	}

	return models.UserID(id), nil
}

func (r *AuthRepository) GetUserID(username string, password string) (models.UserID, error) {

	id, err := r.Storage.GetUser(username, password)
	if err != nil {
		return "", err
	}

	return models.UserID(id), nil
}
