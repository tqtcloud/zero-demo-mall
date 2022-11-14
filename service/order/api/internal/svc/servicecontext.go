package svc

import (
	"github.com/tqtcloud/mall/service/order/api/internal/config"
	"github.com/tqtcloud/mall/service/order/rpc/orderclient"
	"github.com/tqtcloud/mall/service/order/rpc/types/order"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	OrderRpc order.OrderClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		OrderRpc: orderclient.NewOrder(zrpc.MustNewClient(c.OrderRpc)),
	}
}
