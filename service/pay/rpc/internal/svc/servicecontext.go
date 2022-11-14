package svc

import (
	"github.com/tqtcloud/mall/service/order/rpc/orderclient"
	"github.com/tqtcloud/mall/service/order/rpc/types/order"
	"github.com/tqtcloud/mall/service/pay/model"
	"github.com/tqtcloud/mall/service/pay/rpc/internal/config"
	"github.com/tqtcloud/mall/service/user/rpc/types/github.com/tqtcloud/mall/service/user"
	"github.com/tqtcloud/mall/service/user/rpc/userclient"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	PayModel model.PayModel
	UserRpc  user.UserClient
	OrderRpc order.OrderClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:   c,
		PayModel: model.NewPayModel(conn, c.CacheRedis),
		UserRpc:  userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		OrderRpc: orderclient.NewOrder(zrpc.MustNewClient(c.OrderRpc)),
	}
}
