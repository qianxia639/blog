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

	return pageList, nil
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
