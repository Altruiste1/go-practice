package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"go-practice/go_zero/go_im/user/api/internal/config"
	"go-practice/go_zero/go_im/user/api/internal/middleware"
	"go-practice/go_zero/go_im/user/rpc/userclient"
)

type ServiceContext struct {
	Config config.Config
	userclient.User
	LoginVerification rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:            c,
		User:              userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		LoginVerification: middleware.NewLoginVerificationMiddleware().Handle,
	}

}
