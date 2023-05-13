package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

func Md5(data []byte) string {
	return fmt.Sprintf("%x", md5.Sum(data))
}

func ShardMd5(localfile string) (string, error) {
	file, err := os.Open(localfile)
	if err != nil {
		return "", err
	}
	defer file.Close()

	buf := make([]byte, 4086)
	hash := md5.New()
	for {
		n, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}
		hash.Write(buf[:n])
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}
