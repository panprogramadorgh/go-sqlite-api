package utils

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func ConnectDB(filepath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func CheckDB(db *sql.DB) error {
	return db.Ping()
}
