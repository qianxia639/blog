package request

type Post struct {
	UserId      int64  `json:"userId" binding:"required"`
	TypeId      uint32 `json:"typeId" binding:"required"`
	Description string `json:"description" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Content     string `json:"content" binding:"required"`
	Flag        string `json:"flag" binding:"required"`
	Tags        []string
	Selected    []string
}
