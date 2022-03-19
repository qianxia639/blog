package utils

import (
	"crypto/md5"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func Md5(password []byte) string {
	hash := md5.Sum(password)
	return fmt.Sprintf("%x", hash)
}

func Encrypt(password string) (string, error) {
	pwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(pwd), err
}

func Decrypt(password, chechPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(password), []byte(chechPassword))
}
