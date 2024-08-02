package utils

import "golang.org/x/crypto/bcrypt"

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

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
