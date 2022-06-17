package request

type SaveBlog struct {
	TypeId  uint16
	Title   string
	Content string
	Flag    string
	Tags    []string
}

type UpdateBlog struct {
	Id      uint64
	Title   string
	Content string
	Flag    string
}
