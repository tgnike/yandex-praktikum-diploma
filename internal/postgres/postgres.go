package postgres

import "github.com/jackc/pgx"

type Storage struct {
	DB             *pgx.ConnPool
	DataSourceName string
}

func (s *Storage) Init() {

}
