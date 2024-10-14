package seeds

import (
	"go-practice/go_zero/shop/initnal/databases"
	"go-practice/go_zero/shop/model"
)

func SeedsMysql() {
	databases.MysqlDefaultDB.AutoMigrate(model.Spu{})
}
