package api

import (
	"waf-srv/pkg/invoker"
	"waf-srv/pkg/router/core"

	"github.com/hehanpeng/gofund/proto/fund/gen/upssrv"
)

func Test(c *core.Context) {
	reply, err := invoker.UpsSrvGrpc.SayHello(c.Context, &upssrv.HelloRequest{
		Name: "test",
	})
	if err != nil {
		c.JSONE(1, "测试rpc失败", err)
		return
	}
	c.JSONOK(reply)
}
