package response

type Blog struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	UpdatedAt int64  `json:"updatedAt"`
}
