package main

import (
	"flag"
	"fmt"
	"go-practice/go_zero/shop/initnal"
	"go-practice/go_zero/shop/initnal/global"

	"go-practice/go_zero/shop/internal/handler"
	"go-practice/go_zero/shop/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/shop-api.yaml", "the config file")

func main() {
	flag.Parse()

	conf.MustLoad(*configFile, &global.Cfg)

	server := rest.MustNewServer(global.Cfg.RestConf)
	defer server.Stop()

	initnal.Initial()

	ctx := svc.NewServiceContext(*global.Cfg)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", global.Cfg.Host, global.Cfg.Port)
	server.Start()
}
