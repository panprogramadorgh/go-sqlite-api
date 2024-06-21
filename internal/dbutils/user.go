package dbutils

import (
	"database/sql"

	utils "github.com/panprogramadorgh/jsonwebtokenserver/internal/utils"
)

func GetUsers(db *sql.DB) (utils.Users, error) {
	query :=
		`
		SELECT * FROM users ORDER BY user_id DESC
	`
	rows, err := db.Query(query, nil)
	if err != nil {
		return nil, err
	}

	var users utils.Users

	for rows.Next() {
		var userId int
		var username string
		var password string
		var firstname string
		var lastname string
		rows.Scan(&userId, &username, &password, &firstname, &lastname)
		user := utils.User{
			UserID: userId,
			UserPayload: utils.UserPayload{
				Credentials: utils.Credentials{
					Username: username,
					Password: password,
				},
				Firstname: firstname,
				Lastname:  lastname,
			},
		}
		users = append(users, user)
	}

	return users, nil
}

func PostUser(db *sql.DB, user utils.UserPayload) error {
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

func DeleteUsers(db *sql.DB) error {
	query :=
		`
	DELETE FROM users WHERE TRUE
	`
	_, err := db.Exec(query, nil)
	return err
}
