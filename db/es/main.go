package main

import (
	"context"
	"fmt"
	elasticsearch8 "github.com/elastic/go-elasticsearch/v8"
)

var EsClient *elasticsearch8.TypedClient

func main() {
	//s8, _ := elasticsearch8.NewDefaultClient()
	// ES 配置

}
func init() {
	cfg := elasticsearch8.Config{
		Addresses: []string{
			"http://localhost:9200",
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

// updateDocument2 更新文档

func deleteDocument(client *elasticsearch8.TypedClient) {
	resp, err := client.Delete("my-review-1", "1").
		Do(context.Background())
	if err != nil {
		fmt.Printf("delete document failed, err:%v\n", err)
		return
	}
	fmt.Printf("result:%v\n", resp.Result)
}

// deleteIndex 删除 index
func deleteIndex(client *elasticsearch8.TypedClient) {
	resp, err := client.Indices.
		Delete("my-review-1").
		Do(context.Background())
	if err != nil {
		fmt.Printf("delete document failed, err:%v\n", err)
		return
	}
	fmt.Printf("Acknowledged:%v\n", resp.Acknowledged)
}
