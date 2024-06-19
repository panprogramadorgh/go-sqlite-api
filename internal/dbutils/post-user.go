package dbutils

import (
	"database/sql"

	"github.com/panprogramadorgh/jsonwebtokenserver/internal/utils"
)

func PostUser(db *sql.DB, user utils.User) error {
	query :=
		`
	INSERT INTO users (username, password, firstname, lastname) VALUES (?, ?, ?, ?)
	`
	_, err := db.Exec(query, user.Username, user.Password, user.Firstname, user.Lastname)
	if err != nil {
		return err
	}
	return nil
}
