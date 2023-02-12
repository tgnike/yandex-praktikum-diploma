package postgres

import (
	"context"

	"github.com/google/uuid"
)

func (s *Storage) GetUser(ctx context.Context, username string, password string) (string, error) {

	sqlStatement := `SELECT uid from users where username = $1 and password =$2`
	row := s.DB.QueryRow(ctx, sqlStatement, username, password)

	var uid string

	err := row.Scan(&uid)

	if err != nil {
		return "", err
	}

	return uid, nil
}

func (s *Storage) StoreUser(ctx context.Context, username string, password string) (string, error) {

	uid := uuid.New().String()

	sqlStatement := `INSERT INTO users (uid, username, password) VALUES ($1, $2, $3)`
	_, err := s.DB.Exec(ctx, sqlStatement, uid, username, password)

	if err != nil {
		// TODO Добавить типы ошибок
		return "", err
	}

	return uid, nil
}
