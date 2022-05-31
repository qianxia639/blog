package system

import (
	"encoding/json"
	"sync"

	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/model"
	"github.com/qianxia/blog/model/response"
	"github.com/qianxia/blog/utils"
	"gorm.io/gorm"
)

type SearchService struct{}

// 根据title搜索博客
//
// GET /blog/_search
// {
// 	"query": {
// 	  "bool": {
// 		"must": [
// 		  {
// 			"multi_match": {
// 			  "query": "测试",
// 			  "fields": ["title","description"]
// 			}
// 		  }
// 		],
// 		"filter": {
// 		  "term": {
// 			"title": "据"
// 		  }
// 		}
// 	  }
// 	},
// 	"highlight": {
// 	  "pre_tags": "<span style='color:red'>",
// 	  "post_tags": "</span>",
// 	  "fields": {
// 		"description": {},
// 		"title": {}
// 	  }
// 	},
// 	"from": 0,
// 	"size": 5,
// 	"sort": [
// 	  {
// 		"views": {
// 		  "order": "desc"
// 		}
// 	  },
// 	]
//   }
func (*SearchService) SearchBlog(title string, pageNo, pageSize int) (*response.PageList, error) {
	// var pageList response.PageList
	// pageList.Total = total
	// pageList.DataList = blogs

	// query := map[string]interface{}{
	// 	"query": map[string]interface{}{
	// 		"bool": map[string]interface{}{
	// 			"must": map[string]interface{}{
	// 				"multi_match": map[string]interface{}{
	// 					"query":  title,
	// 					"fields": []string{"title", "description"},
	// 				},
	// 			},
	// 			"filter": map[string]interface{}{
	// 				"term": map[string]interface{}{
	// 					"title": string([]rune(title)[0]),
	// 				},
	// 			},
	// 		},
	// 	},
	// 	"highlight": map[string]interface{}{
	// 		"pre_tags":  "<span style='color:#07b9ff'>",
	// 		"post_tags": "</span>",
	// 		"fields": map[string]interface{}{
	// 			"title":       map[string]interface{}{},
	// 			"description": map[string]interface{}{},
	// 		},
	// 	},
	// 	"from": (pageNo - 1) * pageSize,
	// 	"size": pageSize,
	// 	"sort": map[string]interface{}{
	// 		"views": map[string]interface{}{
	// 			"order": "desc",
	// 		},
	// 	},
	// }

	tmpl, err := utils.Loadtemplate(title, pageNo, pageSize)
	if err != nil {
		return nil, err
	}

	res, err := SystemGroups.ElasticSearchService.Search2("blog", tmpl)

	// res, err := SystemGroups.ElasticSearchService.Search("blog", query)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, err
	}

	// 将total和dataList封装到pageList中
	var pageList response.PageList
	pageList.Total = int64(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64))
	var wg sync.WaitGroup
	resp := make([]response.Search, 0, pageList.Total)
	// ch := make(chan []response.Search, pageList.Total)

	// 遍历返回信息中hits的hits
	wg.Add(1)
	go func() {
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
			// ch <- resp
		}
		wg.Done()
	}()
	wg.Wait()
	pageList.PageNo = pageNo
	pageList.PageSize = pageSize
	pageList.DataList = resp

	return &pageList, nil
}

/**
* 根据title和时间进行搜索
 */
func (*SearchService) SearchPriBlog(title, startDate, endDate string, pageSize, pageNo int, userId uint64) (pageList response.PageList, err error) {
	var blogs []response.Blog
	var total int64

	if title != "" && startDate == "" && endDate == "" {
		err = global.QX_DB.Debug().Model(&model.Blog{}).Where("title LIKE ? AND user_id = ?", "%"+title+"%", userId).Offset((pageNo - 1) * pageSize).Limit(pageSize).Find(&blogs).Count(&total).Error
	} else if title == "" && startDate != "" && endDate != "" {
		err = global.QX_DB.Debug().Model(&model.Blog{}).Where("updated_at BETWEEN UNIX_TIMESTAMP(?) AND UNIX_TIMESTAMP(?)", startDate, endDate).Offset((pageNo - 1) * pageSize).Limit(pageSize).Find(&blogs).Count(&total).Error
	} else if title == "" && startDate == "" && endDate == "" {
		err = global.QX_DB.Debug().Model(&model.Blog{}).Where("user_id = ?", userId).Offset((pageNo - 1) * pageSize).Limit(pageSize).Find(&blogs).Count(&total).Error
	} else {
		err = global.QX_DB.Debug().Model(&model.Blog{}).Scopes(func(db *gorm.DB) *gorm.DB {
			return db.Where("title LIKE ? AND user_id = ?", "%"+title+"%", userId)
		}, func(db *gorm.DB) *gorm.DB {
			return db.Where("updated_at BETWEEN UNIX_TIMESTAMP(?) AND UNIX_TIMESTAMP(?)", startDate, endDate)
		}).Offset((pageNo - 1) * pageSize).Limit(pageSize).Find(&blogs).Count(&total).Error
	}

	// global.QX_DB.Model(&model.Blog{}).Where("title LIKE ?", "%"+title+"%").Count(&total)

	// var pageList response.PageList

	pageList.Total = total
	pageList.PageNo = pageNo
	pageList.PageSize = pageSize

	pageList.DataList = blogs

	return
}
