package router

import (
	v1 "ups-srv/api/v1"
	"ups-srv/pkg/invoker"

	"github.com/gotomicro/ego/server/egin"
	"github.com/gotomicro/ego/server/egrpc"
	"github.com/hehanpeng/gofund/proto/fund/gen/upssrv"
)

type Ups struct {
	upssrv.UnimplementedUpsServer
}

func ServeGRPC() *egrpc.Component {
	srv := egrpc.Load("server.grpc").Build()
	upssrv.RegisterUpsServer(srv.Server, &Ups{})
	return srv
}

func GetRouter() *egin.Component {
	Router := invoker.Gin
	// 方便统一添加路由组前缀 多服务器上线使用
	PublicGroup := Router.Group("api")
	{
		PublicGroup.POST("hello", v1.Hello)
	}
	return Router
}
