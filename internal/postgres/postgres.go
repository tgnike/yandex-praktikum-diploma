package postgres

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/google/uuid"

	"github.com/jackc/pgx/v5"
)

type Storage struct {
	DB             *pgx.Conn
	DataSourceName string
}

const migrationURL string = "file://migrations/postgres"

func (s *Storage) Init() {
	conn, err := pgx.Connect(context.Background(), s.DataSourceName)

	if err != nil {
		log.Panicf("postgres init error %v", err)
	}

	s.DB = conn

	err = s.loadMigrations()

	if errors.Is(err, migrate.ErrNoChange) {
		return
	}

	if err != nil {
		log.Fatalf("error migrations %v", err)
	}
}

func (s *Storage) loadMigrations() error {
	m, err := migrate.New(
		migrationURL,
		s.DataSourceName)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil {
		return err
	}

	return nil

}

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

func (s *Storage) GetOrder(ctx context.Context, orderNumber string, userID string) error {

	return nil
}

func (s *Storage) CommitOrder(ctx context.Context, order string, status string, balance float32, userID string) error {

	sqlStatement := `INSERT INTO orders (ordernumber, useruid, balance,status,date) VALUES ($1, $2, $3, $4, $5)`
	_, err := s.DB.Exec(ctx, sqlStatement, order, userID, balance, status, time.Now())

	if err != nil {
		// TODO Добавить типы ошибок
		return err
	}

	return nil

}

func (s *Storage) UpdateOrder(ctx context.Context, order string, status string, balance float32) error {

	sqlStatement := `UPDATE orders set balance = $2, status= $3 WHERE ordernumber= $1 `
	_, err := s.DB.Exec(ctx, sqlStatement, order, balance, status)

	if err != nil {
		// TODO Добавить типы ошибок
		return err
	}

	return nil

}

func (s *Storage) GetUserOrders(ctx context.Context, userID string) error {
	return nil
}
