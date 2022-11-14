package svc

import "github.com/tqtcloud/mall/service/user/rpcgoct/internal/config"

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
