package dbutils

import (
	"database/sql"
	"fmt"

	utils "github.com/panprogramadorgh/jsonwebtokenserver/internal/utils"
)

func GetUsers(db *sql.DB) ([]utils.User, error) {
	query :=
		`
		SELECT * FROM users ORDER BY user_id DESC
	`
	rows, err := db.Query(query, nil)
	if err != nil {
		return nil, err
	}

	var users []utils.User

	for rows.Next() {
		var userId int
		var username string
		var password string
		var firstname string
		var lastname string
		err := rows.Scan(&userId, &username, &password, &firstname, &lastname)
		if err != nil {
			fmt.Printf("error al escanear fila: %v", err)
		}
		user := utils.User{
			UserID: userId,
			Credentials: utils.Credentials{
				Username: username,
				Password: password,
			},
			Firstname: firstname,
			Lastname:  lastname,
		}
		users = append(users, user)
	}

	return users, nil
}
