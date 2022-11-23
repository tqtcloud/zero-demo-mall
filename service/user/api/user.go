package main

import (
	"flag"
	"fmt"
	"github.com/tqtcloud/mall/service/user/api/internal/config"
	"github.com/tqtcloud/mall/service/user/api/internal/handler"
	"github.com/tqtcloud/mall/service/user/api/internal/svc"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	// 自定义错误
	//httpx.SetErrorHandler(func(err error) (int, interface{}) {
	//	switch e := err.(type) {
	//	case *errorx.CodeError:
	//		fmt.Println("命中11111111111111111111111")
	//		return http.StatusOK, *e.Data()
	//	default:
	//		return http.StatusInternalServerError, nil
	//	}
	//})
	// 关闭输出的统计日志(stat)
	logx.DisableStat()
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
