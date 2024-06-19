package dbutils

import "database/sql"

func CheckDB(db *sql.DB) error {
	return db.Ping()
}
