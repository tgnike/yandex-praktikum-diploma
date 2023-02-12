package postgres

import (
	"context"
	"time"

	"github.com/tgnike/yandex-praktikum-diploma/internal/models"
)

func (s *Storage) Withdraw(ctx context.Context, order string, sum float32, user string) error {

	sqlStatement := `INSERT INTO withdrawals (ordernumber, useruid, sum,date) VALUES ($1, $2, $3, $4)`
	_, err := s.DB.Exec(ctx, sqlStatement, order, user, sum, time.Now())

	if err != nil {
		return err
	}

	return nil

}

func (s *Storage) GetBalance(ctx context.Context, user string) (*models.Balance, error) {

	sqlStatement := `SELECT accrual, withdrown from balance where useruid = $1`
	row := s.DB.QueryRow(ctx, sqlStatement, user)

	var left float32
	var used float32

	err := row.Scan(&left, &used)

	if err != nil {
		return nil, err
	}

	return &models.Balance{Current: left, Withdrawn: used}, nil

}

func (s *Storage) Withdrawals(ctx context.Context, user string) ([]*models.Withdrawal, error) {

	sqlStatement := `SELECT ordernumber, sum,date from withdrawals where useruid = $1`
	rows, err := s.DB.Query(ctx, sqlStatement, user)

	ws := make([]*models.Withdrawal, 0)

	for rows.Next() {

		var order string
		var sum float32
		var date time.Time

		err := rows.Scan(&order, &sum, &date)

		if err != nil {
			break
		}

		ws = append(ws, &models.Withdrawal{Order: models.OrderNumber(order), Sum: sum, Date: date})

	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return ws, nil

}
