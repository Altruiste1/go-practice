package initnal

import (
	"go-practice/go_zero/shop/initnal/global"
	"go-practice/go_zero/shop/initnal/seeds"
)

func Initial() {
	global.InitGlobal()
	seeds.SeedsMysql()
}
