package svc

import (
	"api/internal/config"
	"api/internal/middleware"
	"github.com/kiyomi-niunai/user/model"
	"github.com/kiyomi-niunai/user/userclient"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc userclient.User
	UserModel model.UsersModel
	Example rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:  c,
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		UserModel: model.NewUsersModel(conn, c.CacheRedis),
		Example: middleware.NewExampleMiddleware().Handle,
	}
}