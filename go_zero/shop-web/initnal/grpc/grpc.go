package grpc

import (
	"go-practice/go_zero/shop/initnal/global"
	"google.golang.org/grpc"
)

func InitGrpc() {
	var err error
	global.GrpcClient, err = grpc.Dial(global.Cfg.GrpcClient)
	if err != nil {
		panic(err)
	}
}
