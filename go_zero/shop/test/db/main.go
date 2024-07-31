package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/zeromicro/go-zero/rest"
	"go-practice/go_zero/shop/initnal/global"
	"go-practice/go_zero/shop/internal/config"
	"go-practice/go_zero/shop/model"
	"go-practice/go_zero/shop/model/esModel"
	"strings"
)

// 同步
func init() {
	global.Cfg = &config.Config{
		RestConf: rest.RestConf{},
		Mysql: config.MySqlConfig{
			Dsn: "hl:123456@tcp(121.199.46.42:3306)/shop?charset=utf8&parseTime=True&loc=Local",
		},
		ElasticSearch: config.ElasticConfig{
			Address: "http://localhost:9200",
		},
	}
	global.InitGlobal()
}

func main() {
	SyncSpu()
}

func SyncSpu() {
	syncSpu()
}

func Bulk(body string, indexName string) error {
	req := esapi.BulkRequest{
		Index: "spu",
		Body:  strings.NewReader(body),
	}
	res, err := req.Do(context.Background(), global.EsClient)
	if err != nil {
		return err
	}
	if res != nil && res.Body != nil {
		defer res.Body.Close()
	}
	return nil
}

func getTestSpuData() []*esModel.Spu {
	spus := make([]*model.Spu, 0)
	global.DB.Model(model.Spu{}).Find(&spus)
	esSpu := make([]*esModel.Spu, 0)
	for _, v := range spus {
		esSpu = append(esSpu, &esModel.Spu{
			Id:          v.Id,
			GoodsName:   v.GoodsName,
			Category:    v.Category,
			Description: v.Description,
		})
	}
	return esSpu
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

func syncSpu() error {
	return Bulk(getBody(), "spu")
}
