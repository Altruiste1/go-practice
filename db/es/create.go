package main

import (
	"context"
	"fmt"
	elasticsearch8 "github.com/elastic/go-elasticsearch/v8"
	"strconv"
)

// 创建一个名为 my-review-1 的 index。
func createIndex() {
	resp, err := EsClient.Indices.
		Create("spu").
		Do(context.Background())
	if err != nil {
		fmt.Printf("create index failed, err:%v\n", err)
		return
	}
	fmt.Printf("index:%#v\n", resp.Index)
}

// 创建一条 document 并添加到 my-review-1 的 index 中。
// indexDocument 索引文档
func indexDocument(client *elasticsearch8.TypedClient) {
	// 定义 document 结构体对象
	d1 := Spu{
		ID:        1,
		GoodsName: "",
		ShopName:  "",
		Price:     0,
		Content:   "",
		Tag:       nil,
		Status:    0,
		Factory:   "",
	}

	// 添加文档
	resp, err := client.Index("my-review-1").
		Id(strconv.FormatInt(d1.ID, 10)).
		Document(d1).
		Do(context.Background())
	if err != nil {
		fmt.Printf("indexing document failed, err:%v\n", err)
		return
	}
	fmt.Printf("result:%#v\n", resp.Result)
}
