package model

import "time"

// 博客表
type Blog struct {
	Id        uint64    `json:"id,omitempty" gorm:"primaryKey;comment:博客Id"`                // 博客Id
	UserId    uint64    `json:"userId,omitempty" gorm:"not null;comment:用户Id"`              // 用户id
	TypeId    uint16    `json:"typeId,omitempty" gorm:"not null;comment:分类Id"`              // 分类id
	Nickname  string    `json:"nickname,omitempty" gorm:"size:20;not null;comment:用户名"`     // 用户名
	TypeName  string    `json:"typeName,omitempty" gorm:"size:20;not null;comment:分类类名"`    // 分类类名
	Title     string    `json:"title,omitempty" gorm:"size:20;not null;comment:标题"`         // 标题
	Content   string    `json:"content,omitempty" gorm:"type:longtext;not null;comment:内容"` // 内容
	Flag      string    `json:"flag,omitempty" gorm:"type:varchar(10);not null;comment:标记"` // 标记
	Views     uint32    `json:"views,omitempty" gorm:"default:0;comment:浏览次数"`              // 浏览次数
	CreatedAt time.Time `json:"createdAt,omitempty" gorm:"type:timestamp;comment:创建时间"`     // 创建时间
	UpdatedAt time.Time `json:"updatedAt,omitempty" gorm:"type:timestamp;comment:更新时间"`     // 更新时间
	Comments  []Comment `json:"Comments,omitempty"`
	Tags      []Tag     `json:"Tags,omitempty" gorm:"many2many:blog_tag;"`
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
		"content": {
		  "type": "text",
		  "analyzer": "ik_max_word"
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
