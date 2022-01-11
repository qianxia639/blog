package model

type Blog struct {
	// 主键
	Id int64 `binding:"required"`
	// 用户id，
	UserId int64
	// 分类id
	TypeId int
	// 标题
	Title string `json:"title" binding:"required"`
	// 内容
	Content string `json:"content" binding:"required"`
	// 浏览次数
	Views int `json:"views"`
	// 标记
	Flag string `json:"flag"`
	// 是否开启点赞
	Were bool `json:"were"`
	// 点赞数
	Likes int `json:"likes"`
	// 是否显示转载声明
	ShareStatement bool `json:"share_statement"`
	// 是否开启评论
	Commentabled bool `json:"commentabled"`
	// 是否发布
	Published bool `json:"published"`
	// 创建时间
	CreateTime Time `json:"create_time"`
	// 更新时间
	UpdateTime Time `json:"update_time"`
}
