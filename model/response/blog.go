package response

type Blog struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Publish   bool   `json:"publish"`
	UpdatedAt int64  `json:"updatedAt"`
}
