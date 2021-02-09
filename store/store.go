package store

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Store struct {
	DB *sql.DB
}

func New(dataSourceName string) (*Store, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	return &Store{
		DB: db,
	}, nil
}
