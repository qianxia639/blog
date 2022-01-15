package vo

type Post struct {
	UserId   int64  `json:"user_id" binding:"required"`
	TypeId   int    `json:"type_id" binding:"required"`
	Title    string `json:"title" binding:"required"`
	Content  string `json:"content" binding:"required"`
	Flag     string `json:"flag" binding:"required"`
	Tags     []string
	Selected []string
}
