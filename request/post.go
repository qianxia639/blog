package request

type Post struct {
	TypeId      uint8  `json:"typeId"`
	UserId      uint64 `json:"userId"`
	Description string `json:"description"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Flag        string `json:"flag"`
	Tags        []string
	Selected    []string
}
