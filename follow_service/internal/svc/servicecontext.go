package svc

import (
	"context"
	"follow_system/follow_service/internal/config"
)

type ServiceContext struct {
	Config config.Config
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
	return &ServiceContext{
		Config: c,
	}
}
