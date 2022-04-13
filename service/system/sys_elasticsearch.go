package system

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/elastic/go-elasticsearch/v7/esutil"
	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/model"
)

var ElasticSearch = new(elasticSearchService)

type elasticSearchService struct{}

func (e *elasticSearchService) IndicesMapping() error {
	resp, err := global.QX_ES.Indices.Exists([]string{"blog"})
	if err != nil {
		return errors.New(fmt.Sprintf("Error: Indices Exists: %s", err))
	}

	if resp.StatusCode == 404 {
		global.QX_ES.Indices.Create(
			"blog",
			global.QX_ES.Indices.Create.WithBody(strings.NewReader(model.BlogMapping)),
		)
	}

	return err
}

func (e *elasticSearchService) Insert(index, id string, data interface{}) (*esapi.Response, error) {
	// 插入数据到elasticsearch中
	return esapi.IndexRequest{
		Index:      index,
		DocumentID: id,
		Body:       esutil.NewJSONReader(data),
		Refresh:    "true",
	}.Do(context.Background(), global.QX_ES)
}

func (e *elasticSearchService) Delete(index, id string) (*esapi.Response, error) {
	return esapi.DeleteRequest{
		Index:      index,
		DocumentID: id,
	}.Do(context.Background(), global.QX_ES)
}

func (e *elasticSearchService) Update(index, id string, data map[string]interface{}) (*esapi.Response, error) {
	// 修改elasticsearch中对应的文档记录
	return esapi.UpdateRequest{
		Index:      index,
		DocumentID: id,
		Body:       esutil.NewJSONReader(&data),
	}.Do(context.Background(), global.QX_ES)
}

func (e *elasticSearchService) Search(index string, data bytes.Buffer) (*esapi.Response, error) {
	return esapi.SearchRequest{
		Index:          []string{index},
		Body:           &data,
		TrackTotalHits: true,
	}.Do(context.Background(), global.QX_ES)
}
