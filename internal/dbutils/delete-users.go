package dbutils

import "database/sql"

func DeleteUsers(db *sql.DB) error {
	query :=
		`
	DELETE FROM users WHERE TRUE
	`
	_, err := db.Exec(query, nil)
	return err
}
