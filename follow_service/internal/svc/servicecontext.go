package svc

import (
	"context"
	"follow_system/follow_service/internal/config"

	"github.com/go-redis/redis/v8"
)

type ServiceContext struct {
	Config     config.Config
	CacheRedis *redis.Client
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
	return &ServiceContext{
		Config:     c,
		CacheRedis: cacheRedis,
	}
}
