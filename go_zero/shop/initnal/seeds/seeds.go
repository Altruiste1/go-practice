package seeds

import (
	"go-practice/go_zero/shop/initnal/global"
	"go-practice/go_zero/shop/model"
)

func SeedsMysql() {
	global.DB.AutoMigrate(&model.Spu{})
}
