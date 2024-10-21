package svc

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-practice/go_zero/go_im/user/models"
	"go-practice/go_zero/go_im/user/rpc/internal/config"
)

type ServiceContext struct {
	Config    config.Config
	Cache     cache.CacheConf
	UserModel models.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.Dsn)
	return &ServiceContext{
		Config:    c,
		UserModel: models.NewUserModel(sqlConn, c.Cache),
	}
}
