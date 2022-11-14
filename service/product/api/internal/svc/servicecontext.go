package svc

import (
	"github.com/tqtcloud/mall/service/product/api/internal/config"
	"github.com/tqtcloud/mall/service/product/rpc/productclient"
	"github.com/tqtcloud/mall/service/product/rpc/types/product"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	ProductRpc product.ProductClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		ProductRpc: productclient.NewProduct(zrpc.MustNewClient(c.ProductRpc)),
	}
}
