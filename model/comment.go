package model

type Comment struct {
	// 主键
	Id int64 `json:"id" binding:"required"`
	// 博客id
	BlogId int64 `json:"blog_id"`
	// 评论id
	ParentCommentId int64 `json:"parent_comment_id"`
	// 用户名
	Username string `json:"username" binding:"required"`
	// 头像
	Avatar string `json:"avatar"`
	// 评论内容
	Comment string `json:"comment" binding:"required"`
	// 评论时间
	CommentTime Time `json:"comment_time"`
}
