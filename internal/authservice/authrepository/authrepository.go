package authrepository

import (
	"context"

	"github.com/tgnike/yandex-praktikum-diploma/internal/models"
)

type AuthRepository struct {
	Storage Storage
}

type Storage interface {
	CommitUser(ctx context.Context, username string, password string) (string, error)
	GetUser(ctx context.Context, username string, password string) (string, error)
}

func New(storage Storage) *AuthRepository {
	return &AuthRepository{Storage: storage}
}

func (r *AuthRepository) CommitUser(ctx context.Context, username string, password string) (models.UserID, error) {

	id, err := r.Storage.CommitUser(ctx, username, password)
	if err != nil {
		return "", err
	}

	return models.UserID(id), nil
}

func (r *AuthRepository) GetUserID(ctx context.Context, username string, password string) (models.UserID, error) {

	id, err := r.Storage.GetUser(ctx, username, password)
	if err != nil {
		return "", err
	}

	return models.UserID(id), nil
}
