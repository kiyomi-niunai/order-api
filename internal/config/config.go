package config

import (
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/rest"
)
import "github.com/tal-tech/go-zero/zrpc"

type Config struct {
	rest.RestConf
	UserRpc zrpc.RpcClientConf
	Mysql struct {
		DataSource string
	}
	Auth      struct {
		AccessSecret string
			AccessExpire int64
	}
	CacheRedis cache.ClusterConf
}
