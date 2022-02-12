package model

type Blog struct {
	// 主键
	Id int64 `json:"id"`
	// 用户id，
	UserId int64 `json:"userId" gorm:"NOT NULL"`
	// 分类id
	TypeId uint8 `json:"typeId" gorm:"NOT NULL"`
	// 标题
	Title string `json:"title" gorm:"size:20;NOT NULL"`
	// 描述
	Description string `json:"description" gorm:"NOT NULL"`
	// 内容
	Content string `json:"content" gorm:"type:text;NOT NULL"`
	// 标记
	Flag string `json:"flag" gorm:"type:varchar(10);comment:标记"`
	// 是否开启点赞
	Were bool `json:"were" gorm:"comment:是否开启点赞"`
	// 是否显示转载声明
	ShareStatement bool `json:"shareStatement" gorm:"comment:是否显示转载声明"`
	// 是否开启评论
	EnableComment bool `json:"enableComment" gorm:"comment:是否开启评论"`
	// 点赞数
	Likes uint32 `json:"likes" gorm:"comment:点赞数"`
	// 浏览次数
	Views uint32 `json:"views" gorm:"comment:浏览次数"`
	// 创建时间
	CreatedAt int64 `json:"createdAt" gorm:"autoCreateTine:milli;comment:创建时间"`
	// 更新时间
	UpdatedAt int64 `json:"updatedAt" gorm:"autoUpdateTine:milli;comment:更新时间"`
	Tags      []Tag `gorm:"many2many:blog_tag;"`
}
