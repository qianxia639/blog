package request

type Post struct {
	Id          uint64 `json:"id"`
	TypeId      uint16 `json:"typeId"`
	UserId      uint64 `json:"userId"`
	Description string `json:"description"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Flag        string `json:"flag"`
	Tags        []string
}
