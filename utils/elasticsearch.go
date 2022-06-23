package utils

import (
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/qianxia/blog/global"
)

func ElasticSearch() *elasticsearch.Client {

	client, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: global.QX_CONFIG.ElasticSearch.Addr,
	})
	if err != nil {
		global.QX_LOG.Fatalf("ElasticSearch Connection Error: %v", err)
		return nil
	}

	return client
}
