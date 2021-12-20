package utils

import (
	"strings"

	"github.com/google/uuid"
)

func CreateUUID() string {
	u, err := uuid.NewUUID()
	if err != nil {
		panic(err)
	}
	return strings.ReplaceAll(u.String(), "-", "")
}
