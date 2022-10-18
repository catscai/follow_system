package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/gateway"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Gateway     gateway.GatewayConf
	CacheRedis  redis.RedisConf
	FollowMysql struct {
		DataSource string
	}
}
