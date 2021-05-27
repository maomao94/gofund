package router

import (
	"github.com/gotomicro/ego/server/egrpc"
)

//type Resource struct {
//	resourcesrv.UnimplementedResourceServer
//}

func ServeGRPC() *egrpc.Component {
	srv := egrpc.Load("server.grpc").Build()
	//resourcesrv.RegisterResourceServer(srv.Server, &Resource{})
	return srv
}
