package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"strings"
)

func Bulk(body string, indexName string) error {
	req := esapi.BulkRequest{
		Index: "spu",
		Body:  strings.NewReader(body),
	}
	res, err := req.Do(context.Background(), EsClient)
	if err != nil {
		return err
	}
	if res != nil && res.Body != nil {
		defer res.Body.Close()
	}
	return nil
}

func getTestSpuData() []Spu {
	return []Spu{
		{
			ID:        0,
			GoodsName: "纳爱斯牙膏",
			ShopName:  "天猫",
			Price:     12,
			Content:   "",
			Tag:       nil,
			Status:    0,
			Factory:   "纳爱斯",
		},
		{
			ID:        0,
			GoodsName: "大宝SOD蜜",
			ShopName:  "广州宝洁官方旗舰店",
			Price:     10,
			Content:   "",
			Tag:       nil,
			Status:    0,
			Factory:   "广州宝洁",
		},
	}
}
func getBody() string {
	s := ""
	for _, v := range getTestSpuData() {
		s = s + fmt.Sprintf("{ \"index\" : { \"_index\" : \"%s\" } }\n", "spu")
		b, _ := json.Marshal(v)
		s = s + string(b) + "\n"
	}
	fmt.Printf("%s", s)
	return s
}

// { "index" : { "_index" : "library" } }
// {"name":"萧山区图书馆","address":"杭州市萧山区"}
// { "index" : { "_index" : "library" } }
// {"name":"余杭区图书馆","address":"杭州市余杭区"}
func syncSpu() error {
	return Bulk(getBody(), "spu")
}
