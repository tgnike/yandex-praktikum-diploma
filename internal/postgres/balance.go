package postgres

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/tgnike/yandex-praktikum-diploma/internal/models"
)

func getBalance(ctx context.Context, tx pgx.Tx, user string) (float32, float32, error) {

	sqlBalance := `SELECT accrual, withdrown FROM balance WHERE useruid=$1`
	row := tx.QueryRow(ctx, sqlBalance, user)

	var currentBalance float32
	var withdrown float32

	err := row.Scan(&currentBalance, &withdrown)

	if err != nil {
		return 0, 0, err
	}

	return currentBalance, withdrown, nil

}

func updateBalance(ctx context.Context, tx pgx.Tx, user string, accrual float32, withdwown float32) error {

	sqlBalanceUpdate := `INSERT INTO balance (useruid, accrual, withdrown) VALUES ( $1, $2, $3) 
	ON CONFLICT (useruid) DO UPDATE
	SET accrual = $2, withdrown = $3`
	_, err := tx.Exec(ctx, sqlBalanceUpdate, user, accrual, withdwown)

	if err != nil {
		return err
	}

	return nil

}

func (s *Storage) Withdraw(ctx context.Context, order string, sum float32, user string) error {

	tx, err := s.DB.BeginTx(ctx, TxDefOpts)

	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)

	currentBalance, withdrown, err := getBalance(ctx, tx, user)

	if err != nil {
		return err
	}

	if sum > currentBalance {
		return errors.New("not enough accruals")
	}

	sqlWithdrawals := `INSERT INTO withdrawals (ordernumber, useruid, sum,date) VALUES ($1, $2, $3, $4)`
	_, err = tx.Exec(ctx, sqlWithdrawals, order, user, sum, time.Now())

	if err != nil {
		return err
	}

	err = updateBalance(ctx, tx, user, currentBalance-sum, withdrown+sum)

	if err != nil {
		return err
	}

	tx.Commit(ctx)

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

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ws, nil
		} else {
			return nil, err
		}
	}

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
