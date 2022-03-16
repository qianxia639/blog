package response

type Blog struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Views     uint32 `json:"views"`
	UpdatedAt int64  `json:"updatedAt"`
}
