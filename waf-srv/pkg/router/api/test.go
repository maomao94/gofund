package api

import (
	"waf-srv/model/response"
	"waf-srv/pkg/invoker"
	"waf-srv/pkg/router/core"

	"google.golang.org/grpc/status"

	"github.com/hehanpeng/gofund/proto/fund/gen/upssrv"
)

func Test(c *core.Context) {
	reply, err := invoker.UpsSrvGrpc.SayHello(c.Context, &upssrv.HelloRequest{
		Name: "test",
	})
	if err != nil {
		response.FailWithMessage(status.Code(err).String(), c.Context)
		return
	}
	response.OkWithData(reply, c.Context)
}
