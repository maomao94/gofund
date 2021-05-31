package router

import (
	"waf-srv/pkg/invoker"
	"waf-srv/pkg/router/api"
	"waf-srv/pkg/router/core"

	"github.com/gotomicro/ego/server/egin"
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
func GetRouter() *egin.Component {
	r := invoker.Gin
	r.GET("/api/test", core.Handle(api.Test))
	return r
}
