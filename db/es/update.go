package main

import (
	"context"
	"encoding/json"
	"fmt"
	elasticsearch8 "github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/update"
)

// updateDocument 更新文档
func updateDocument(client *elasticsearch8.TypedClient) {
	// 修改后的结构体变量
	d1 := Spu{
		ID:        0,
		GoodsName: "",
		ShopName:  "",
		Price:     0,
		Content:   "",
		Tag:       nil,
		Status:    0,
		Factory:   "",
	}

	resp, err := client.Update("my-review-1", "1").
		Doc(d1). // 使用结构体变量更新
		Do(context.Background())
	if err != nil {
		fmt.Printf("update document failed, err:%v\n", err)
		return
	}
	fmt.Printf("result:%v\n", resp.Result)
}

func updateDocument2(client *elasticsearch8.TypedClient) {
	// 修改后的JSON字符串
	str := `{
					"id":1,
					"userID":147982601,
					"score":5,
					"content":"这是一个二次修改后的好评！",
					"tags":[
						{
							"code":1000,
							"title":"好评"
						},
						{
							"code":9000,
							"title":"有图"
						}
					],
					"status":2,
					"publishDate":"2023-12-10T15:27:18.219385+08:00"
				}`
	// 直接使用JSON字符串更新
	resp, err := client.Update("my-review-1", "1").
		Request(&update.Request{
			Doc: json.RawMessage(str),
		}).
		Do(context.Background())
	if err != nil {
		fmt.Printf("update document failed, err:%v\n", err)
		return
	}
	fmt.Printf("result:%v\n", resp.Result)
}
