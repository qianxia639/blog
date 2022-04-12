package system

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/model"
)

var ElasticSearch = new(elasticSearchService)

type elasticSearchService struct{}

func (e *elasticSearchService) IndicesAndMapping() error {
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
	var buf bytes.Buffer
	if err := e.jsonEncode(&buf, data); err != nil {
		return nil, err
	}

	// 插入数据到elasticsearch中
	return esapi.IndexRequest{
		Index:      index,
		DocumentID: id,
		Body:       bytes.NewReader(buf.Bytes()),
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
	var buf bytes.Buffer
	if err := e.jsonEncode(&buf, data); err != nil {
		return nil, err
	}
	// 修改elasticsearch中对应的文档记录
	return esapi.UpdateRequest{
		Index:      index,
		DocumentID: id,
		Body:       bytes.NewReader(buf.Bytes()),
	}.Do(context.Background(), global.QX_ES)
}

func (e *elasticSearchService) Search(index string, data map[string]interface{}) (*esapi.Response, error) {
	var buf bytes.Buffer
	if err := e.jsonEncode(&buf, data); err != nil {
		return nil, err
	}

	return esapi.SearchRequest{
		Index:          []string{index},
		Body:           bytes.NewReader(buf.Bytes()),
		TrackTotalHits: true,
	}.Do(context.Background(), global.QX_ES)
}

func (e *elasticSearchService) jsonEncode(buf io.Writer, data interface{}) error {
	return json.NewEncoder(buf).Encode(data)
}
