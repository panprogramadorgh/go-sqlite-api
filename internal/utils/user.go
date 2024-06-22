package utils

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

type Credentials struct {
	Username string
	Password string
}

type UserPayload struct {
	Credentials
	Firstname string
	Lastname  string
}

type User struct {
	UserID int
	UserPayload
}

type Users []User

func Auth(db *sql.DB, username string, password string) bool {
	user, err := GetUser(db, username)
	if err != nil || user == nil {
		return false
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return false
	}
	return true
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
// func (u *Users) ReloadData(db *sql.DB) error {
// 	users, err := GetUsers(db)
// 	if err != nil {
// 		return nil
// 	}
// 	*u = users
// 	return nil
// }

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
		var userID int
		var username string
		var password string
		var firstname string
		var lastname string
		err := rows.Scan(&userID, &username, &password, &firstname, &lastname)
		if err != nil {
			return nil, err
		}
		p := User{
			UserID: userID,
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

func GetUser(db *sql.DB, username string) (*User, error) {
	query :=
		`
	SELECT * FROM users WHERE username = ?
	`
	rows, err := db.Query(query, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var user *User = nil
	for rows.Next() {
		var userID int
		var username string
		var password string
		var firstname string
		var lastname string
		err := rows.Scan(&userID, &username, &password, &firstname, &lastname)
		if err != nil {
			return nil, err
		}
		user = &User{
			UserID: userID,
			UserPayload: UserPayload{
				Credentials: Credentials{
					Username: username,
					Password: password,
				},
				Firstname: firstname,
				Lastname:  lastname,
			},
		}
	}
	return user, nil
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

func DeleteUser(db *sql.DB, username string) error {
	// Eliminar filas con el username del payload
	if _, err := db.Exec(`
	DELETE FROM users WHERE username = ?
	`, username); err != nil {
		return err
	}
	return nil
}
