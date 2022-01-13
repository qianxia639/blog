package blog

import (
	"github.com/jinzhu/gorm"
	"github.com/qianxia/blog/model"
	"github.com/qianxia/blog/utils"
)

type BlogService struct {
	DB *gorm.DB
}

func NewBlogService() BlogService {
	return BlogService{DB: utils.GetDB()}
}

func (bs BlogService) Save(blog model.Blog) (uint, error) {
	var err error

	// 构建数据
	// newBlog := model.Blog{
	// 	Id:             utils.NextId(),
	// 	UserId:         1,
	// 	TypeId:         1,
	// 	Title:          blog.Title,
	// 	Content:        blog.Content,
	// 	Flag:           blog.Flag,
	// 	Were:           blog.Were,
	// 	ShareStatement: blog.ShareStatement,
	// 	Commentabled:   blog.Commentabled,
	// 	CreateTime:     model.Time(time.Now()),
	// 	UpdateTime:     model.Time(time.Now()),
	// }

	return 0, err
}
