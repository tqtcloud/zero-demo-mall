package svc

import (
	"github.com/tqtcloud/mall/service/order/model"
	"github.com/tqtcloud/mall/service/order/rpc/internal/config"
	"github.com/tqtcloud/mall/service/product/rpc/productclient"
	"github.com/tqtcloud/mall/service/product/rpc/types/product"
	"github.com/tqtcloud/mall/service/user/rpc/types/github.com/tqtcloud/mall/service/user"
	"github.com/tqtcloud/mall/service/user/rpc/userclient"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	OrderModel model.OrderModel

	UserRpc    user.UserClient
	ProductRpc product.ProductClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:     c,
		OrderModel: model.NewOrderModel(conn, c.CacheRedis),
		UserRpc:    userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		ProductRpc: productclient.NewProduct(zrpc.MustNewClient(c.ProductRpc)),
	}
}
