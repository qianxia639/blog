package system

import (
	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/model"
	"github.com/qianxia/blog/model/response"
)

type SearchService struct{}

// @function SearchBlog
// @description 根据title搜索博客
// @param title string, pageNo, pageSize int
// @return *response.PageList, error
func (s *SearchService) SearchBlog(title string, pageNo, pageSize int) (*response.PageList, error) {

	var blogs []model.Blog
	var total int64
	offset := (pageNo - 1) * pageSize
	err := global.DB.Debug().Preload("Tags").Preload("User").Where("title LIKE ?", "%"+title+"%").Limit(pageSize).Offset(offset).Find(&blogs).Count(&total).Error

	pageList := &response.PageList{
		Total:    total,
		PageNo:   pageNo,
		PageSize: pageSize,
		DataList: blogs,
	}
	return pageList, err
}
