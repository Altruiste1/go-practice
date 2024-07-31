package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Mysql         MySqlConfig
	ElasticSearch ElasticConfig
}

type MySqlConfig struct {
	Dsn string
}

type ElasticConfig struct {
	Address string
}
