package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

type Payload struct {
	ID       string    `json:"id"`
	Username string    `json:"username"`
	IssuedAt time.Time `json:"issuee_at"`
	ExpireAt time.Time `json:"expire_at"`
}

func NewPayload(userename string, duration time.Duration) *Payload {
	return &Payload{
		ID:       uuid.NewString(),
		Username: userename,
		IssuedAt: time.Now(),
		ExpireAt: time.Now().Add(duration),
	}
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpireAt) {
		return ErrExpiredToken
	}

	return nil
}
