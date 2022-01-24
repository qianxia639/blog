package utils

import (
	"github.com/google/uuid"
)

func CreateUUID() string {
	u, err := uuid.NewUUID()
	if err != nil {
		panic(err)
	}
	return u.String()
}
