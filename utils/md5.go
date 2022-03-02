package utils

import (
	"crypto/md5"
	"fmt"
)

func Md5(password []byte) string {
	hash := md5.Sum(password)
	return fmt.Sprintf("%x", hash)
}
