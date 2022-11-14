package svc

import (
	"github.com/tqtcloud/mall/service/pay/api/internal/config"
	"github.com/tqtcloud/mall/service/pay/rpc/payclient"
	"github.com/tqtcloud/mall/service/pay/rpc/types/pay"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	PayRpc pay.PayClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		PayRpc: payclient.NewPay(zrpc.MustNewClient(c.PayRpc)),
	}
}
