package system

import (
	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/model"
	"github.com/qianxia/blog/model/response"
	"gorm.io/gorm"
)

type SearchService struct{}

/**
* 根据title或description搜索博客
 */
func (*SearchService) SearchBlog(query string) (*response.PageList, error) {
	var (
		// 获取total
		total int64
		blogs []model.Blog
		// 获取dataList
		// blogs []response.Index
	)
	if err := global.QX_DB.Debug().Select("id,user_id,type_id,username,type_name,title,description,updated_at").Preload("Tags").Where("title LIKE ? OR description LIKE ?", "%"+query+"%", "%"+query+"%").Find(&blogs).Count(&total).Error; err != nil {
		return nil, err
	}

	// for _, v := range b {
	// 	var users model.User
	// 	if err := global.QX_DB.Debug().Select("username,avatar").Where("id = ?", v.UserId).Find(&users).Error; err != nil {
	// 		return nil, err
	// 	}
	// 	var types model.Type
	// 	if err := global.QX_DB.Debug().Select("type_name").Where("id = ?", v.TypeId).Find(&types).Error; err != nil {
	// 		return nil, err
	// 	}

	// 	blogs = append(blogs, response.Index{
	// 		Id:          v.Id,
	// 		Title:       v.Title,
	// 		Description: v.Description,
	// 		UpdatedAt:   utils.TimestampToString(v.UpdatedAt),
	// 		TypeName:    types.TypeName,
	// 		Avatar:      users.Avatar,
	// 		Username:    users.Username,
	// 		Tags:        v.Tags,
	// 	})
	// }

	// 将total和dataList封装到pageList中
	var pageList response.PageList
	pageList.Total = total
	pageList.DataList = blogs

	return &pageList, nil
}

/**
* 根据title和时间进行搜索
 */
func (*SearchService) SearchPriBlog(title, startDate, endDate string, pageSize, pageNum int, userId uint64) (pageList response.PageList, err error) {
	var blogs []response.Blog
	var total int64

	if title == "" && startDate == "" && endDate == "" {
		err = global.QX_DB.Debug().Model(&model.Blog{}).Where("user_id = ?", userId).Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&blogs).Count(&total).Error

	} else {
		err = global.QX_DB.Debug().Model(&model.Blog{}).Scopes(func(db *gorm.DB) *gorm.DB {
			return db.Where("title LIKE ? AND user_id = ?", "%"+title+"%", userId)
		}, func(db *gorm.DB) *gorm.DB {
			return db.Where("updated_at BETWEEN UNIX_TIMESTAMP(?) AND UNIX_TIMESTAMP(?)", startDate, endDate)
		}).Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&blogs).Count(&total).Error
	}

	// global.QX_DB.Model(&model.Blog{}).Where("title LIKE ?", "%"+title+"%").Count(&total)

	// var pageList response.PageList

	pageList.Total = total
	pageList.PageNum = pageNum
	pageList.PageSize = pageSize

	pageList.DataList = blogs

	return
}
