package model

type Comment struct {
	// 主键
	Id int64 `binding:"required"`
	// 博客id
	BlogId int64
	// 评论id
	ParentCommentId int64
	// 用户名
	Username string `json:"username" binding:"required"`
	// 头像
	Avatar string `json:"avatar"`
	// 评论内容
	Comment string `json:"comment" binding:"required"`
	// 评论时间
	CommentTime Time `json:"comment_time"`
}
