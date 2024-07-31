package main

import (
	"context"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"testing"
)

func TestBulk(t *testing.T) {
	t.Log(syncSpu())
	resp, err := EsClient.Search().
		Index("spu").
		Query(&types.Query{
			MatchAll: &types.MatchAllQuery{},
		}).
		Do(context.Background())
	if err != nil {
		fmt.Printf("search document failed, err:%v\n", err)
		return
	}
	fmt.Printf("total: %d\n", resp.Hits.Total.Value)
	// 遍历所有结果
	for _, hit := range resp.Hits.Hits {
		fmt.Printf("%s\n", hit.Source_)
	}

}
