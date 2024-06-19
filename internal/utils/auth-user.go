package utils

func (u Users) AuthUser(username string, password string) bool {
	userI := u.IndexOfUserPerUsername(username)
	if userI == -1 {
		return false
	}
	user := u[userI]
	auth := user.Credentials.Username == username && user.Credentials.Password == password
	return auth
}
