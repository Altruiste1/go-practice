package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql Mysql
	Cache cache.CacheConf
}

type Cache struct {
	Host string
	Type string
	Pass string
}

type Mysql struct {
	Dsn string
}
