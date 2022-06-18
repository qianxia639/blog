package request

type Comment struct {
	BlogId   uint64 `binding:"required"`
	ParentId uint64
	Nickname string `default:"匿名用户" binding:"required"`
	Content  string `binding:"required"`
}
