package request

type ParentConment struct {
	BlogId   uint64
	Nickname string `json:"nickname" default:"匿名用户"`
	Content  string
}

type ChildComment struct {
	BlogId   uint64
	ParentId uint64
	Nickname string `json:"nickname" default:"匿名用户"`
	Content  string
}
