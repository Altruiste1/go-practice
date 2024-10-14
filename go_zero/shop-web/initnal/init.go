package initnal

import (
	"go-practice/go_zero/shop/initnal/databases"
	"go-practice/go_zero/shop/initnal/grpc"
	"go-practice/go_zero/shop/initnal/seeds"
)

func Initial() {
	databases.InitDatabases()
	seeds.SeedsMysql()
	grpc.InitGrpc()
}
