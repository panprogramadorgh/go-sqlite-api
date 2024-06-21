package utils

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
