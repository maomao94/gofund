package router

import (
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
