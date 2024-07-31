package global

import (
	"fmt"
	elasticsearch8 "github.com/elastic/go-elasticsearch/v8"
)

func initElasticSearch() {
	cfg := elasticsearch8.Config{
		Addresses: []string{
			Cfg.ElasticSearch.Address,
		},
	}

	var err error
	// 创建客户端连接
	EsClient, err = elasticsearch8.NewTypedClient(cfg)
	if err != nil {
		fmt.Printf("elasticsearch.NewTypedClient failed, err:%v\n", err)
		return
	}
}
