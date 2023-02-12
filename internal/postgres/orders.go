package postgres

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/tgnike/yandex-praktikum-diploma/internal/models"
	"github.com/tgnike/yandex-praktikum-diploma/internal/server"
)

func (s *Storage) CommitOrder(ctx context.Context, order string, status string, balance float32, userID string) error {

	tx, err := s.DB.BeginTx(ctx, TxDefOpts)

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

func (s *Storage) UpdateOrder(ctx context.Context, order string, status string, balance float32, user string) error {

	tx, err := s.DB.BeginTx(ctx, TxDefOpts)
	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)

	currentBalance, withdrown, err := getBalance(ctx, tx, user)

	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return err
	}

	sqlStatementOrder := `UPDATE orders set balance = $2, status= $3 WHERE ordernumber= $1 `
	_, err = tx.Exec(ctx, sqlStatementOrder, order, balance, status)

	if err != nil {
		// TODO Добавить типы ошибок
		return err
	}

	err = updateBalance(ctx, tx, user, currentBalance+balance, withdrown)

	if err != nil {
		return err
	}

	tx.Commit(ctx)

	return nil

}
func (s *Storage) GetOrder(ctx context.Context, orderNumber string, userID string) error {
	return nil
}

func (s *Storage) GetUserOrders(ctx context.Context, userID string, orders models.OrderContainerInterface) error {
	sqlStatement := `SELECT ordernumber, balance, status, date from orders where useruid = $1 ORDER BY date asc`
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
