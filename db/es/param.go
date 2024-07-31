package main

// 定义与 document 数据对应的 Review 和 Tag 结构体，分别表示"评价"和"评价标签"。
type Spu struct {
	ID        int64   `json:"id"`
	GoodsName string  `json:"goodsName"`
	ShopName  string  `json:"shopName"`
	Price     float64 `json:"price"`
	Content   string  `json:"content"`
	Tag       []Tag   `json:"tag"`
	Status    int     `json:"status"`
	Factory   string  `json:"factory"`
}

type Tag struct {
}
