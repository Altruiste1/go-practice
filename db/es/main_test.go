package main

import (
	"context"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/combinedfieldsoperator"
	"testing"
)

func TestBulk(t *testing.T) {
	//t.Log(syncSpu())
	resp, err := EsClient.Search().
		Index("spu").
		Query(&types.Query{
			CombinedFields: &types.CombinedFieldsQuery{
				AutoGenerateSynonymsPhraseQuery: nil,
				Boost:                           nil,
				Fields:                          []string{"Description", "GoodsName"},
				MinimumShouldMatch:              nil,
				Operator:                        nil,
				Query:                           "搬运工",
				QueryName_:                      nil,
				ZeroTermsQuery:                  nil,
			},
			//Match: map[string]types.MatchQuery{
			//	"GoodsName": {
			//		Query: "农夫",
			//	},
			//	//"Description": {
			//	//	Query: "大自然",
			//	//},
			//},
			//MatchAll: &types.MatchAllQuery{
			//	QueryName_: nil,
			//},
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
func TestMultiMatch(t *testing.T) {
	multiMatch("spu", []string{"GoodsName", "Description"}, "搬运工")
	return
}

func TestCombineMatch(t *testing.T) {
	combineMatch("spu", []string{"GoodsName", "Description"}, "搬运工")
	return
}

func combineMatch(index string, fields []string, val string) {
	resp, err := EsClient.Search().
		Index(index).
		Query(&types.Query{
			CombinedFields: &types.CombinedFieldsQuery{
				AutoGenerateSynonymsPhraseQuery: nil,
				Boost:                           nil,
				Fields:                          fields,
				MinimumShouldMatch:              nil,
				Operator:                        &combinedfieldsoperator.And,
				Query:                           val,
				QueryName_:                      nil,
				ZeroTermsQuery:                  nil,
			},
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
func multiMatch(index string, fields []string, val string) {
	resp, err := EsClient.Search().
		Index(index).
		Query(&types.Query{
			MultiMatch: &types.MultiMatchQuery{
				AutoGenerateSynonymsPhraseQuery: nil,
				Boost:                           nil,
				Fields:                          fields,
				MinimumShouldMatch:              nil,
				Operator:                        nil,
				Query:                           val,
				QueryName_:                      nil,
				ZeroTermsQuery:                  nil,
			},
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
