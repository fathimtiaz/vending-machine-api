package sql

import "database/sql"

type SQLRepository struct {
	db *sql.DB
}

func NewSQLRepository(driver, connStr string) (repo *SQLRepository, err error) {
	repo.db, err = sql.Open(driver, connStr)
	return
}
