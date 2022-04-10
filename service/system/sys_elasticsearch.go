package system

import (
	"strings"

	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/model"
)

var ElasticSearch = new(elasticSearch)

type elasticSearch struct{}

func (e *elasticSearch) IndicesAndMapping() {
	resp, err := global.QX_ES.Indices.Exists([]string{"blog"})
	if err != nil {
		global.QX_LOG.Errorf("Error: Indices.Exists: %s", err)
	}

	if resp.StatusCode == 404 {
		global.QX_ES.Indices.Create(
			"blog",
			global.QX_ES.Indices.Create.WithBody(strings.NewReader(model.BlogMapping)),
		)
	}
}
