package utils

type Credentials struct {
	Username string
	Password string
}

type User struct {
	UserID int
	Credentials
	Firstname string
	Lastname  string
}

type Users []User
