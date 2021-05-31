package router

import (
	"github.com/gotomicro/ego/server/egin"
	"github.com/gotomicro/ego/server/egrpc"
	"github.com/hehanpeng/gofund/proto/fund/gen/upssrv"
	"ups-srv/pkg/invoker"
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
	r := invoker.Gin
	//r.GET("/api/test",api.Test)
	return r
}
