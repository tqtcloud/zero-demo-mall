package svc

import (
	"github.com/tqtcloud/mall/service/user/api/internal/config"
	"github.com/tqtcloud/mall/service/user/rpc/types/github.com/tqtcloud/mall/service/user"
	"github.com/tqtcloud/mall/service/user/rpc/userclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc user.UserClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
