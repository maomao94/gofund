package v1

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gotomicro/ego/core/etrace"
	"github.com/hehanpeng/gofund/common/resp"
	"waf-srv/pkg/invoker"
)

func Test(c *gin.Context) {
	span, ctx := etrace.StartSpanFromContext(context.Background(), "callHTTP()")
	defer span.Finish()

	req := invoker.UpsHttpComp.R()
	// Inject traceId Into Header
	c1 := etrace.HeaderInjector(ctx, req.Header)

	info, err := req.SetContext(c1).Get("/api/hello")
	if err != nil {
		resp.Fail(c)
		return
	}
	resp.OkWithData(info.String(), c)
}

func Hello(c *gin.Context) {
	resp.Ok(c)
}
