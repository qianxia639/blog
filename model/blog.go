package model

type Blog struct {
	// 主键
	Id uint64 `json:"id" gorm:"primaryKey"`
	// 用户id，
	UserId uint64 `json:"userId" gorm:"NOT NULL;comment:用户Id"`
	// 分类id
	TypeId uint16 `json:"typeId" gorm:"NOT NULL;comment:分类Id"`
	// 标题
	Title string `json:"title" gorm:"size:20;NOT NULL;comment:标题"`
	// 描述
	Description string `json:"description" gorm:"NOT NULL;comment:描述"`
	// 内容
	Content string `json:"content" gorm:"type:longtext;NOT NULL;comment:内容"`
	// 标记
	Flag string `json:"flag" gorm:"type:varchar(10);comment:标记"`
	// 浏览次数
	Views uint32 `json:"views" gorm:"comment:浏览次数"`
	// 是否已发布
	Publish bool `json:"publish" gorm:"comment:是否已发布"`
	// 创建时间
	CreatedAt int64 `json:"createdAt" gorm:"autoCreateTine:milli;comment:创建时间"`
	// 更新时间
	UpdatedAt int64 `json:"updatedAt" gorm:"autoUpdateTine:milli;comment:更新时间"`
	Tags      []Tag `gorm:"many2many:blog_tag;"`
}
