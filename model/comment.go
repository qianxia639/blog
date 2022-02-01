package model

type Comment struct {
	// 主键
	Id int64 `json:"id" binding:"required"`
	// 博客id
	BlogId int64 `json:"blogId" binding:"required" gorm:"NOT NULL"`
	// 评论id
	ParentCommentId int64 `json:"parentCommentId"`
	// 用户名
	Username string `json:"usParentername" binding:"required" gorm:"size:20;NOT NULL"`
	// 头像
	Avatar string `json:"avatar" binding:"required" gorm:"NOT NULL"`
	// 评论内容
	Comment string `json:"comment" binding:"required" gorm:"NOT NULL"`
	// 评论时间
	// CreatedAt Time `json:"createdAt" gorm:"type:timestamp"`
	CreatedAt Time `json:"createdAt"`
}
