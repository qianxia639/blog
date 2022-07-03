package model

// 博客表
type Blog struct {
	Id     uint64 `json:"id,omitempty" gorm:"primaryKey;comment:博客Id"`   // 博客Id
	UserId uint64 `json:"userId,omitempty" gorm:"not null;comment:用户Id"` // 用户id
	TypeId uint16 `json:"typeId,omitempty" gorm:"not null;comment:分类Id"` // 分类id
	// Nickname string `json:"nickname,omitempty" gorm:"size:20;not null;comment:用户昵称"` // 用户昵称
	User      User      `json:"User,omitempty" gorm:"foreignkey:UserId"`
	TypeName  string    `json:"typeName,omitempty" gorm:"size:20;not null;comment:分类类名"`    // 分类类名
	Title     string    `json:"title,omitempty" gorm:"size:20;not null;comment:标题"`         // 标题
	Content   string    `json:"content,omitempty" gorm:"type:longtext;not null;comment:内容"` // 内容
	Flag      string    `json:"flag,omitempty" gorm:"type:varchar(10);not null;comment:标记"` // 标记
	Views     uint32    `json:"views,omitempty" gorm:"default:0;comment:浏览次数"`              // 浏览次数
	CreatedAt int64     `json:"createdAt,omitempty" gorm:"autoCreateTime,comment:创建时间"`     // 创建时间
	UpdatedAt int64     `json:"updatedAt,omitempty" gorm:"autoCreateTime,comment:更新时间"`     // 更新时间
	Comments  []Comment `json:"Comments,omitempty"`
	Tags      []Tag     `json:"Tags,omitempty" gorm:"many2many:t_blog_tag"`
}

func (*Blog) TableName() string {
	return "t_blog"
}
