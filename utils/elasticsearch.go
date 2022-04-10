package utils

import (
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/qianxia/blog/global"
)

func ElasticSearch() *elasticsearch.Client {
	client, err := elasticsearch.NewDefaultClient()
	if err != nil {
		global.QX_LOG.Fatalf("ElasticSearch Connection Error: %v", err)
		return nil
	}

	return client
}
