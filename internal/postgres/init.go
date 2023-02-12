package postgres

import (
	"context"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var TxDefOpts = pgx.TxOptions{IsoLevel: pgx.ReadCommitted, AccessMode: pgx.ReadWrite}

type Storage struct {
	DB             *pgxpool.Pool
	DataSourceName string
}

const migrationURL string = "file://migrations/postgres"

func (s *Storage) Init() {
	conn, err := pgxpool.New(context.Background(), s.DataSourceName)

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

	// withdrawals
	_, err = s.DB.Exec(context.Background(), `CREATE TABLE IF NOT EXISTS withdrawals (
		id int generated always as identity ( cache 10 ) primary key,
		ordernumber varchar(20) not null 
		, useruid varchar(36) not null
		, sum float not null
		, date TIMESTAMP WITH TIME ZONE
	 )`)

	if err != nil {
		log.Panicf("postgres withdrawals error %v", err)
	}

	// balance
	_, err = s.DB.Exec(context.Background(), `CREATE TABLE IF NOT EXISTS balance (
		useruid varchar(36) primary key
		, accrual float not null
		, withdrown float not null
	 )`)

	if err != nil {
		log.Panicf("postgres balance error %v", err)
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
