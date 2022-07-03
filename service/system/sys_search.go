package system

import (
	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/model"
	"github.com/qianxia/blog/model/response"
)

type SearchService struct{}

// 根据title搜索博客
func (s *SearchService) SearchBlog(title string, pageNo, pageSize int) (*response.PageList, error) {

	var blogs []model.Blog
	var total int64
	offset := (pageNo - 1) * pageSize
	err := global.DB.Debug().Preload("Tags").Preload("User").Where("title LIKE ?", "%"+title+"%").Limit(pageSize).Offset(offset).Find(&blogs).Count(&total).Error

	// result := make([]response.BlogResult, 0, len(blogs))
	// var user model.User
	// for _, v := range blogs {
	// 	err := global.DB.Debug().Model(user).Where("nickname = ?", v.Nickname).First(&user).Error
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	result = append(result, response.BlogResult{
	// 		Id:        v.Id,
	// 		Nickname:  v.Nickname,
	// 		TypeName:  v.TypeName,
	// 		Title:     v.Title,
	// 		Content:   v.Content,
	// 		Avatar:    user.Avatar,
	// 		Flag:      v.Flag,
	// 		Views:     v.Views,
	// 		UpdatedAt: v.UpdatedAt,
	// 		Tags:      v.Tags,
	// 	})
	// }

	pageList := &response.PageList{
		Total:    total,
		PageNo:   pageNo,
		PageSize: pageSize,
		DataList: blogs,
	}
	return pageList, err
}
