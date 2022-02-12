package app

import (
	"errors"

	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/model"
	"github.com/qianxia/blog/response"
)

type TypeService struct {
}

// 查詢type列表，按amount降序排列
func (ts TypeService) ListOrderByAmountDesc() ([]model.Type, error) {
	var types []model.Type
	if err := global.RY_DB.Debug().Select("id,type_name,amount").Order("amount DESC").Find(&types).Error; err != nil {
		return nil, errors.New("查询失败")
	}
	return types, nil
}

// 只显示分类列表不排序
func (ts TypeService) List() ([]model.Type, error) {
	var types []model.Type
	if err := global.RY_DB.Debug().Select("id,type_name,amount").Find(&types).Error; err != nil {
		return nil, errors.New("查询失败")
	}
	return types, nil
}

func (ts TypeService) typeList(id int) ([]response.Index, error) {
	var blogs []response.Index
	// if err := global.RY_DB.Raw(`SELECT b.id,b.title,b.content,b.update_time,t.type_name,u.avatar,u.username
	// 					FROM t_blog b JOIN t_user u ON u.id = b.user_id JOIN t_type t ON b.type_id = t.id AND b.type_id = ?`, id).Scan(&blogs).Error; err != nil {
	// 	return nil, errors.New("查询失败")
	// }

	// for k, v := range blogs {
	// 	if err := global.RY_DB.Raw(`select t.id,t.tag_name from t_tag t JOIN
	// 				(select DISTINCT(bt.tag_id) from t_blog_tag bt JOIN t_blog b ON bt.blog_id = ?) as tag
	// 				ON t.id = tag.tag_id`, v.Id).Scan(&blogs[k].Tags).Error; err != nil {
	// 		return nil, errors.New("查询失败")
	// 	}
	// }

	return blogs, nil
}
