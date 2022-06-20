package system

import (
	"bytes"
	"encoding/json"

	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/model/command"
	"github.com/qianxia/blog/model/response"
)

type SearchService struct {
	buf bytes.Buffer
}

// 根据title搜索博客
func (s *SearchService) SearchBlog(title string, pageNo, pageSize int) (*response.PageList, error) {
	// var pageList response.PageList
	// pageList.Total = total
	// pageList.DataList = blogs

	// query := map[string]interface{}{
	// 	"query": map[string]interface{}{
	// 		"bool": map[string]interface{}{
	// 			"must": map[string]interface{}{
	// 				"multi_match": map[string]interface{}{
	// 					"query":  title,
	// 					"fields": []string{"title", "content"},
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
	// 			"content": 	   map[string]interface{}{},
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

	buf, err := TemplateServices.SearchBlogByTitle(title, pageNo, pageSize)
	if err != nil {
		global.QX_LOG.Errorf("Error read template file: %s", err)
		return nil, err
	}

	res, err := ElasticSearchServices.Search("blog", buf)
	if err != nil {
		global.QX_LOG.Errorf("Error getting response: %s", err)
		return nil, err
	}

	defer res.Body.Close()

	// var r map[string]interface{}
	var elasticsearchConfig command.ElasticsearchConfig
	if err := json.NewDecoder(res.Body).Decode(&elasticsearchConfig); err != nil {
		global.QX_LOG.Errorf("Error parsing the response body: %s", err)
		return nil, err
	}

	// log.Printf("elasticsearch resposne %v\n", es)

	// 将total和dataList封装到pageList中
	// total := int64(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64))

	// resp := make([]response.Search, 0, total)

	// // 遍历返回信息中hits的hits
	// for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
	// 	var title interface{}
	// 	var content interface{}

	// 	if hit.(map[string]interface{})["highlight"].(map[string]interface{})["title"] == nil {
	// 		title = hit.(map[string]interface{})["_source"].(map[string]interface{})["title"]
	// 	} else {
	// 		title = hit.(map[string]interface{})["highlight"].(map[string]interface{})["title"].([]interface{})[0]
	// 	}

	// 	if hit.(map[string]interface{})["highlight"].(map[string]interface{})["content"] == nil {
	// 		content = hit.(map[string]interface{})["_source"].(map[string]interface{})["content"]
	// 	} else {
	// 		content = hit.(map[string]interface{})["highlight"].(map[string]interface{})["content"].([]interface{})[0]
	// 	}

	// 	resp = append(resp, response.Search{
	// 		Id:        hit.(map[string]interface{})["_source"].(map[string]interface{})["id"],
	// 		UserId:    hit.(map[string]interface{})["_source"].(map[string]interface{})["userId"],
	// 		TypeId:    hit.(map[string]interface{})["_source"].(map[string]interface{})["typeId"],
	// 		TypeName:  hit.(map[string]interface{})["_source"].(map[string]interface{})["typeName"].(string),
	// 		Username:  hit.(map[string]interface{})["_source"].(map[string]interface{})["typeName"].(string),
	// 		Title:     title,
	// 		Content:   content,
	// 		UpdatedAt: hit.(map[string]interface{})["_source"].(map[string]interface{})["updatedAt"],
	// 		Tags:      hit.(map[string]interface{})["_source"].(map[string]interface{})["Tags"],
	// 	})
	// }

	total := elasticsearchConfig.Hits.Total.Value

	resp, err := s.eachHits(total, elasticsearchConfig.Hits.Hits)
	if err != nil {
		global.QX_LOG.Errorf("Error each Hits: %s", err)
		return nil, err
	}

	pageList := s.result(total, pageNo, pageSize, resp)

	return pageList, nil
}

/**
* 根据title和时间进行搜索
 */
func (s *SearchService) SearchPriBlog(title, startDate, endDate string, pageSize, pageNo int, userId uint64) (*response.PageList, error) {
	// var blogs []response.Blog
	// var total int64
	var err error

	// if title != "" && startDate == "" && endDate == "" {
	// 	err = global.QX_DB.Debug().Model(&model.Blog{}).Where("title LIKE ? AND user_id = ?", "%"+title+"%", userId).Offset((pageNo - 1) * pageSize).Limit(pageSize).Find(&blogs).Count(&total).Error
	// } else if title == "" && startDate != "" && endDate != "" {
	// 	err = global.QX_DB.Debug().Model(&model.Blog{}).Where("updated_at BETWEEN ? AND ?", startDate, endDate).Offset((pageNo - 1) * pageSize).Limit(pageSize).Find(&blogs).Count(&total).Error
	// } else if title == "" && startDate == "" && endDate == "" {
	// 	err = global.QX_DB.Debug().Model(&model.Blog{}).Where("user_id = ?", userId).Offset((pageNo - 1) * pageSize).Limit(pageSize).Find(&blogs).Count(&total).Error
	// } else {
	// 	err = global.QX_DB.Debug().Model(&model.Blog{}).Scopes(func(db *gorm.DB) *gorm.DB {
	// 		return db.Where("title LIKE ? AND user_id = ?", "%"+title+"%", userId)
	// 	}, func(db *gorm.DB) *gorm.DB {
	// 		return db.Where("updated_at BETWEEN ? AND ?", startDate, endDate)
	// 	}).Offset((pageNo - 1) * pageSize).Limit(pageSize).Find(&blogs).Count(&total).Error
	// }

	if title != "" && startDate != "" && endDate != "" {
		s.buf, err = TemplateServices.SearchBlogByTitleAndTime(title, startDate, endDate, pageSize, pageNo, userId)
		if err != nil {
			global.QX_LOG.Errorf("Error read template file: %s", err)
			return nil, err
		}
	} else {
		s.buf, err = TemplateServices.SearchBlogByTitleOrTime(title, startDate, endDate, pageSize, pageNo, userId)
		if err != nil {
			global.QX_LOG.Errorf("Error read template file: %s", err)
			return nil, err
		}
	}

	res, err := ElasticSearchServices.Search("blog", s.buf)
	if err != nil {
		global.QX_LOG.Errorf("Error search: %s", err)
		return nil, err
	}
	defer res.Body.Close()

	var elasticsearchConfig command.ElasticsearchConfig
	if err := json.NewDecoder(res.Body).Decode(&elasticsearchConfig); err != nil {
		global.QX_LOG.Errorf("Error parsing the response body: %s", err)
		return nil, err
	}

	total := elasticsearchConfig.Hits.Total.Value

	resp, err := s.eachHits(total, elasticsearchConfig.Hits.Hits)
	if err != nil {
		global.QX_LOG.Errorf("Error each Hits: %s", err)
		return nil, err
	}

	pageList := s.result(total, pageNo, pageSize, resp)

	return pageList, err
}

func (*SearchService) result(total int64, pageNo, pageSize int, data interface{}) *response.PageList {
	return &response.PageList{
		Total:    total,
		PageNo:   pageNo,
		PageSize: pageSize,
		DataList: data,
	}
}

func (s *SearchService) eachHits(total int64, hits []command.Hit) ([]response.Search, error) {
	resp := make([]response.Search, 0, total)
	for _, hit := range hits {
		var content interface{}
		var title interface{}

		if hit.Highlight["content"] == nil {
			content = hit.Source["content"]
		} else {
			content = hit.Highlight["content"]
		}

		if hit.Highlight["title"] == nil {
			title = hit.Source["title"]
		} else {
			title = hit.Highlight["title"]
		}

		err := json.NewEncoder(&s.buf).Encode(hit.Source)
		if err != nil {
			return nil, err
		}

		var blog response.Search
		err = json.NewDecoder(&s.buf).Decode(&blog)
		if err != nil {
			return nil, err
		}
		resp = append(resp, response.Search{
			Id:        blog.Id,
			UserId:    blog.UserId,
			TypeId:    blog.TypeId,
			Title:     title,
			Content:   content,
			Nickname:  blog.Nickname,
			TypeName:  blog.TypeName,
			UpdatedAt: blog.UpdatedAt,
			Tags:      blog.Tags,
		})
	}
	return resp, nil
}
