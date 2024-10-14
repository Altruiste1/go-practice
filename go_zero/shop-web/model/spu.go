package model

type Spu struct {
	BaseModel
	GoodsName   string //商品名
	Category    string //分类
	Description string
	Price       float64
}
