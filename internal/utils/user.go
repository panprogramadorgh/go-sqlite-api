package utils

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

func (u Users) IndexOf(username string) int {
	userI := -1
	for i, eachUser := range u {
		if eachUser.Credentials.Username == username {
			userI = i
			break
		}
	}
	return userI
}

func (u Users) Auth(username string, password string) bool {
	userI := u.IndexOf(username)
	if userI == -1 {
		return false
	}
	user := u[userI]
	usernameMatch := user.Credentials.Username == username
	if !usernameMatch {
		return false
	}
	passwordMatch := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return passwordMatch == nil
}

// Hashea la contraseña del payload del usuario
func (p *UserPayload) HashPass() error {
	pass, err := bcrypt.GenerateFromPassword([]byte(p.Password), 16)
	if err != nil {
		return err
	}
	p.Password = string(pass)
	return nil
}

// Se actualiza el slice de usuarios (tipo subyacente)
func (u *Users) ReloadData(db *sql.DB) error {
	users, err := GetUsers(db)
	if err != nil {
		return nil
	}
	*u = users
	return nil
}

func GetUsers(db *sql.DB) (Users, error) {
	query :=
		`
		SELECT * FROM users ORDER BY user_id DESC
	`
	rows, err := db.Query(query, nil)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users Users

	for rows.Next() {
		var userId int
		var username string
		var password string
		var firstname string
		var lastname string
		rows.Scan(&userId, &username, &password, &firstname, &lastname)
		p := User{
			UserID: userId,
			UserPayload: UserPayload{
				Credentials: Credentials{
					Username: username,
					Password: password,
				},
				Firstname: firstname,
				Lastname:  lastname,
			},
		}
		users = append(users, p)
	}

	return users, nil
}

func PostUser(db *sql.DB, p UserPayload) error {
	query :=
		`
	INSERT INTO users (username, password, firstname, lastname) VALUES (?, ?, ?, ?)
	`
	_, err := db.Exec(query, p.Username, p.Password, p.Firstname, p.Lastname)
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
