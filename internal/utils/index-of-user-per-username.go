package utils

func (u Users) IndexOfUserPerUsername(username string) int {
	userI := -1
	for i, eachUser := range u {
		if eachUser.Credentials.Username == username {
			userI = i
			break
		}
	}
	return userI
}
