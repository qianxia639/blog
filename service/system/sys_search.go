package system

import (
	"encoding/json"

	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/model"
	"github.com/qianxia/blog/model/response"
	"github.com/qianxia/blog/utils"
	"gorm.io/gorm"
)

type SearchService struct{}

/**
* 根据title或description搜索博客
 */
func (*SearchService) SearchBlog(key string) (*response.PageList, error) {
	// var (
	// 	// 获取total
	// 	total int64
	// 	blogs []model.Blog
	// 	// 获取dataList
	// 	// blogs []response.Index
	// )
	// if err := global.QX_DB.Debug().Select("id,user_id,type_id,username,type_name,title,description,updated_at").Preload("Tags").Where("title LIKE ? OR description LIKE ?", "%"+key+"%", "%"+key+"%").Find(&blogs).Count(&total).Error; err != nil {
	// 	return nil, err
	// }

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":  key,
				"fields": []string{"title", "description"},
			},
		},
		"highlight": map[string]interface{}{
			"pre_tags":  "<span style='color:#07b9ff'>",
			"post_tags": "</span>",
			"fields": map[string]interface{}{
				"title":       map[string]interface{}{},
				"description": map[string]interface{}{},
			},
		},
	}

	// var buf bytes.Buffer
	// if err := json.NewEncoder(&buf).Encode(query); err != nil {
	// 	return nil, err
	// }

	// res, err := global.QX_ES.Search(
	// 	global.QX_ES.Search.WithContext(context.Background()),
	// 	global.QX_ES.Search.WithIndex("blog"),
	// 	global.QX_ES.Search.WithBody(&buf),
	// 	global.QX_ES.Search.WithTrackTotalHits(true),
	// )
	res, err := ElasticSearch.Search("blog", query)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, err
	}

	resp := make([]response.Search, 0)
	// 将total和dataList封装到pageList中
	var pageList response.PageList
	pageList.Total = int64(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64))

	// 遍历返回信息中hits的hits
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		var title interface{}
		var description interface{}

		if hit.(map[string]interface{})["highlight"].(map[string]interface{})["title"] == nil {
			title = hit.(map[string]interface{})["_source"].(map[string]interface{})["title"]
		} else {
			title = hit.(map[string]interface{})["highlight"].(map[string]interface{})["title"].([]interface{})[0]
		}

		if hit.(map[string]interface{})["highlight"].(map[string]interface{})["description"] == nil {
			description = hit.(map[string]interface{})["_source"].(map[string]interface{})["description"]
		} else {
			description = hit.(map[string]interface{})["highlight"].(map[string]interface{})["description"].([]interface{})[0]
		}

		resp = append(resp, response.Search{
			Id:          hit.(map[string]interface{})["_source"].(map[string]interface{})["id"],
			UserId:      hit.(map[string]interface{})["_source"].(map[string]interface{})["userId"],
			TypeId:      hit.(map[string]interface{})["_source"].(map[string]interface{})["typeId"],
			TypeName:    hit.(map[string]interface{})["_source"].(map[string]interface{})["typeName"].(string),
			Username:    hit.(map[string]interface{})["_source"].(map[string]interface{})["typeName"].(string),
			Title:       title,
			Description: description,
			UpdatedAt:   utils.TimestampToString(int64(hit.(map[string]interface{})["_source"].(map[string]interface{})["updatedAt"].(float64))),
			Tags:        hit.(map[string]interface{})["_source"].(map[string]interface{})["Tags"],
		})
	}

	pageList.DataList = resp

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
