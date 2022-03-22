package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func Encrypt(password string) (string, error) {
	pwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(pwd), err
}

func Decrypt(password, chechPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(password), []byte(chechPassword))
}
