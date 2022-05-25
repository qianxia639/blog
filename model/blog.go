package model

// 博客表
type Blog struct {
	Id          uint64 `json:"id,omitempty" gorm:"primaryKey;comment:博客Id"`                  // 博客Id
	UserId      uint64 `json:"userId,omitempty" gorm:"comment:用户Id"`                         // 用户id
	TypeId      uint16 `json:"typeId,omitempty" gorm:"comment:分类Id"`                         // 分类id
	Username    string `json:"username,omitempty" gorm:"size:20;comment:用户名"`                // 用户名
	TypeName    string `json:"typeName,omitempty" gorm:"size:20;comment:分类类名"`               // 分类类名
	Title       string `json:"title,omitempty" gorm:"size:20;comment:标题"`                    // 标题
	Description string `json:"description,omitempty" gorm:"comment:描述"`                      // 描述
	Content     string `json:"content,omitempty" gorm:"type:longtext;comment:内容"`            // 内容
	Flag        string `json:"flag,omitempty" gorm:"type:varchar(10);comment:标记"`            // 标记
	Views       uint32 `json:"views,omitempty" gorm:"default:0;comment:浏览次数"`                // 浏览次数
	CreatedAt   int64  `json:"createdAt,omitempty" gorm:"autoCreateTine:milli;comment:创建时间"` // 创建时间
	UpdatedAt   int64  `json:"updatedAt,omitempty" gorm:"autoUpdateTine:milli;comment:更新时间"` // 更新时间
	Tags        []Tag  `gorm:"many2many:blog_tag;"`
}

var BlogMapping = `
{
	"mappings": {
	  "properties": {
		"id": {
		  "type": "long"
		},
		"userId": {
		  "type": "long"
		},
		"typeId": {
		  "type": "integer"
		},
		"username": {
		  "type": "keyword"
		},
		"typeName": {
		  "type": "keyword"
		},
		"title": {
		  "type": "text",
		  "analyzer": "ik_max_word"
		},
		"description": {
		  "type": "text",
		   "analyzer": "ik_max_word"
		},
		"content": {
		  "type": "text"
		},
		"flag": {
		  "type": "keyword"
		},
		"views": {
		  "type": "integer"
		},
		"createdAt": {
		  "type": "date"
		},
		"updatedAt": {
		  "type": "date"
		}
	  }
	}
  }
`
