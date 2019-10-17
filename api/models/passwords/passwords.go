package passwords

import (
	"golang.org/x/crypto/bcrypt"
)

const (
	salt = 10
)

func Encrypt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), salt)
	return string(hash), err
}

func Compare(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false
	}
	return true
}
