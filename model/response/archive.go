package response

type Archive struct {
	Id        uint64 `json:"id"`
	Title     string `json:"title"`
	Flag      string `json:"flag"`
	UpdatedAt uint64 `json:"updatedAt"`
}
