package router

import (
	"github.com/gotomicro/ego/server/egin"
	"waf-srv/api/v1"
	"waf-srv/pkg/invoker"
)

func GetRouter() *egin.Component {
	r := invoker.Gin
	r.GET("/api/test", v1.Test)
	return r
}
