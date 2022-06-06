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

type ElasticSearchService struct{}

var ElasticSearchServices = new(ElasticSearchService)

func (e *ElasticSearchService) IndicesMapping() error {
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
	return nil
}

// 在索引中创建一条文档
func (e *ElasticSearchService) Insert(index, id string, data interface{}) (*esapi.Response, error) {
	return esapi.CreateRequest{
		Index:      index,
		DocumentID: id,
		Body:       esutil.NewJSONReader(data),
		Refresh:    "true",
	}.Do(context.Background(), global.QX_ES)
}

// 在索引中删除一条文档
func (e *ElasticSearchService) Delete(index, id string) (*esapi.Response, error) {
	return esapi.DeleteRequest{
		Index:      index,
		DocumentID: id,
	}.Do(context.Background(), global.QX_ES)
}

func (e *ElasticSearchService) Update(index, id string, data map[string]interface{}) (*esapi.Response, error) {
	// 修改elasticsearch中对应的文档记录
	return esapi.UpdateRequest{
		Index:      index,
		DocumentID: id,
		Body:       esutil.NewJSONReader(&data),
	}.Do(context.Background(), global.QX_ES)
}

func (e *ElasticSearchService) Search(index string, data map[string]interface{}) (*esapi.Response, error) {
	return esapi.SearchRequest{
		Index:          []string{index},
		Body:           esutil.NewJSONReader(&data),
		TrackTotalHits: true,
	}.Do(context.Background(), global.QX_ES)
}

func (e *ElasticSearchService) Search2(index string, data bytes.Buffer) (*esapi.Response, error) {
	return esapi.SearchRequest{
		Index:          []string{index},
		Body:           &data,
		TrackTotalHits: true,
	}.Do(context.Background(), global.QX_ES)
}
