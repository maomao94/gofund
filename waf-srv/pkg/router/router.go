package router

import (
	"waf-srv/pkg/invoker"
	"waf-srv/pkg/router/api"
	"waf-srv/pkg/router/core"

	"github.com/gotomicro/ego/server/egin"
)

func GetRouter() *egin.Component {
	r := invoker.Gin
	r.GET("/api/test", core.Handle(api.Test))
	return r
}
