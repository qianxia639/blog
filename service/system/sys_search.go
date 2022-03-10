package system

import (
	"errors"

	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/model"
	"github.com/qianxia/blog/model/response"
	"github.com/qianxia/blog/utils"
)

type SearchService struct{}

/**
* 根据title和description搜索博客
 */
func (*SearchService) SearchBlog(query string) (*response.PageList, error) {
	var (
		// 获取total
		total int64
		b     []model.Blog
		// 获取dataList
		blogs []response.Index
	)
	if err := global.QX_DB.Debug().Select("id,user_id,type_id,title,description,updated_at").Preload("Tags").Where("title LIKE ? OR description LIKE ?", "%"+query+"%", "%"+query+"%").Find(&b).Count(&total).Error; err != nil {
		global.QX_LOG.Error(err)
		return nil, errors.New("查询失败")
	}

	for _, v := range b {
		var users model.User
		if err := global.QX_DB.Debug().Select("username,avatar").Where("id = ?", v.UserId).Find(&users).Error; err != nil {
			global.QX_LOG.Error(err)
			return nil, errors.New("查询失败")
		}
		var types model.Type
		if err := global.QX_DB.Debug().Select("type_name").Where("id = ?", v.TypeId).Find(&types).Error; err != nil {
			global.QX_LOG.Error(err)
			return nil, errors.New("查询失败")
		}

		blogs = append(blogs, response.Index{
			Id:          v.Id,
			Title:       v.Title,
			Description: v.Description,
			UpdatedAt:   utils.TimestampToString(v.UpdatedAt),
			TypeName:    types.TypeName,
			Avatar:      users.Avatar,
			Username:    users.Username,
			Tags:        v.Tags,
		})
	}

	// 将total和dataList封装到pageList中
	var pageList response.PageList
	pageList.Pagination.Total = total
	pageList.DataList = blogs

	return &pageList, nil
}

/**
* 根据title和时间进行搜索
 */
func (*SearchService) SearchPriBlog(title, startDate, endDate string, pageSize, pageNo int) (*response.PageList, error) {
	var blogs []response.Blog
	var total int64
	// TODO 这里是有问题的
	if err := global.QX_DB.Debug().Model(&model.Blog{}).Select("id,title,publish,updated_at").Where("title LIKE ?", "%"+title+"%").
		Or("updated_at BETWEEN UNIX_TIMESTAMP(?) AND UNIX_TIMESTAMP(?)", startDate, endDate).Offset((pageNo - 1) * pageSize).Limit(pageSize).Find(&blogs).Error; err != nil {
		return nil, err
	}

	global.QX_DB.Model(&model.Blog{}).Where("title LIKE ?", "%"+title+"%").Count(&total)

	var pageList response.PageList

	pageList.Pagination.Total = total
	pageList.Pagination.CurrentPage = pageNo
	pageList.Pagination.PerPage = pageSize
	if int(total)/pageSize == 0 {
		pageList.Pagination.LastPage = int(total) / pageSize
	} else {
		pageList.Pagination.LastPage = int(total)/pageSize + 1
	}

	pageList.DataList = blogs

	return &pageList, nil
}
