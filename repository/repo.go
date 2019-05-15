package repository

import (
	"database/sql"

	"github.com/sirupsen/logrus"
)

// Repo -
type Repo struct{
	db *sql.DB

}

// New -
func New() *Repo {

	// Opening a driver typically will not attempt to connect to the database.
	db, err := sql.Open("driver-name", "database=test1")
	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		logrus.Fatal("invalid connection")
	}
	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(50)

	return &Repo{
		db: db,
	}
}

// GetAccounts -
func (r *Repo) GetAccounts(account string) (map[string]string, error) {
	resp := map[string]string{}

	row := r.db.QueryRow("SELECT account_id FROM accounts WHERE account_id = ? ", account)

	err := row.Scan(resp)
	if err != nil {
		logrus.Error(err, "query fail")
	}

	return resp, nil
}
