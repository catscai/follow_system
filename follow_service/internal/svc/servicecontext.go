package svc

import (
	"context"
	"follow_system/follow_service/internal/config"

	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config      config.Config
	CacheRedis  *redis.Client
	FollowMysql sqlx.SqlConn
}

type CommonData struct {
	ID    int32  `json:"id,omitempty"`
	Email string `json:"email,omitempty"`
}

type TokenContext struct {
	context.Context
	COMM CommonData
}

func NewServiceContext(c config.Config) *ServiceContext {
	cacheRedis := redis.NewClient(&redis.Options{
		Addr:     c.CacheRedis.Host,
		Password: c.CacheRedis.Pass,
		DB:       0,
	})
	conn := sqlx.NewSqlConn("mysql", c.FollowMysql.DataSource)
	return &ServiceContext{
		Config:      c,
		CacheRedis:  cacheRedis,
		FollowMysql: conn,
	}
}
