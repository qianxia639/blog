package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func Encrypt(password string) ([]byte, error) {

	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func Decrypt(currentPassword, targetPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(currentPassword), []byte(targetPassword))
}
