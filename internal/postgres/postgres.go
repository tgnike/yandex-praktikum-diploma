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
	"github.com/tgnike/yandex-praktikum-diploma/internal/models"
	"github.com/tgnike/yandex-praktikum-diploma/internal/server"

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

	_, err = s.DB.Exec(context.Background(), `CREATE TABLE IF NOT EXISTS users (
		id int generated always as identity ( cache 10 ) primary key
		, uid varchar(36) not null unique
		, username text not null unique
		, password text not null )`)

	if err != nil {
		log.Panicf("postgres users error %v", err)
	}

	_, err = s.DB.Exec(context.Background(), `CREATE TABLE IF NOT EXISTS orders (
    
			ordernumber varchar(20) not null primary key
			, useruid varchar(36) not null
			, balance float not null
			, status varchar(25) not null
			 , date TIMESTAMP WITH TIME ZONE
		 )`)

	if err != nil {
		log.Panicf("postgres orders error %v", err)
	}

	// err = s.loadMigrations()

	// if errors.Is(err, migrate.ErrNoChange) {
	// 	return
	// }

	// if err != nil {
	// 	log.Fatalf("error migrations %v", err)
	// }
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

	tx, err := s.DB.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.ReadCommitted, AccessMode: pgx.ReadWrite})

	if err != nil {
		return err
	}

	sqlStatementSelect := `SELECT useruid from orders WHERE ordernumber = $1`

	row := tx.QueryRow(ctx, sqlStatementSelect, order)

	var uid string

	err = row.Scan(&uid)

	// Есть ошибка, но и запрос не пустой
	norows := false
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			norows = true
		} else {
			return err
		}
	}

	if !norows {
		if uid == userID {
			return server.ErrUploadedByUser
		}
		if uid != userID {
			return server.ErrUploadedByOtherUser
		}
	}

	sqlStatementInsert := `INSERT INTO orders (ordernumber, useruid, balance,status,date) VALUES ($1, $2, $3, $4, $5)`
	_, err = tx.Exec(ctx, sqlStatementInsert, order, userID, balance, status, time.Now())

	if err != nil {

		return err
	}

	err = tx.Commit(ctx)

	if err != nil {
		tx.Rollback(ctx)
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

func (s *Storage) GetUserOrders(ctx context.Context, userID string, orders models.OrderContainerInterface) error {
	sqlStatement := `SELECT ordernumber, balance, status, date from orders where useruid = $1`
	rows, err := s.DB.Query(ctx, sqlStatement, userID)

	if err != nil {
		return err
	}

	for rows.Next() {

		var order string
		var balance float32
		var status string
		var date time.Time

		err := rows.Scan(&order, &balance, &status, &date)

		if err != nil {
			break
		}

		orders.Add(order, balance, status, date)

	}

	err = rows.Err()

	if err != nil {
		return err
	}

	return nil

}
