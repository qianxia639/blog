package response

import "time"

type Archive struct {
	Id        uint64    `json:"id"`
	Title     string    `json:"title"`
	Flag      string    `json:"flag"`
	UpdatedAt time.Time `json:"updatedAt"`
}
