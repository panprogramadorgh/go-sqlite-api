package utils

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
	auth := user.Credentials.Username == username && user.Credentials.Password == password
	return auth
}
