package model

import (
	"time"
)

type Leave struct {
	Id        uint64    `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"CreatedAt" gorm:"type:timestamp"`
}
