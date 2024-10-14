package global

import (
	"go-practice/go_zero/shop/internal/config"
	"google.golang.org/grpc"
)

var (
	Cfg        *config.Config
	GrpcClient *grpc.ClientConn
)
