package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func Enbcrypt(password string) ([]byte, error) {

	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func Debcrypt(currentPassword, targetPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(currentPassword), []byte(targetPassword))
}
