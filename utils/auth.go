package utils

// Predefined user credentials (for demonstration purposes)
var Users = map[string]string{
	"user1": "pass1",
	"user2": "pass2",
}

func Authenticate(username, password string) bool {
	if pass, ok := Users[username]; ok && pass == password {
		return true
	}

	return false
}
