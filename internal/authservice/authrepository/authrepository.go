package authrepository

type AuthRepository struct {
	storage Storage
}

type Storage interface {
	SaveUser(username string, password string) (string, error)
	GetUser(username string, password string) (string, error)
}

func New(storage Storage) *AuthRepository {
	return &AuthRepository{storage: storage}
}
