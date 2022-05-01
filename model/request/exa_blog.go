package request

type Post struct {
	Id          uint64 `json:"id"`
	TypeId      uint16 `json:"typeId" binding:"required"`
	UserId      uint64 `json:"userId"`
	Username    string `json:"username"`
	Description string `json:"description" validate:"min=10,max=100"`
	Title       string `json:"title" validate:"max=15,min=1"`
	Content     string `json:"content" binding:"required"`
	Flag        string `json:"flag" binding:"required"`
	Tags        []string
}
