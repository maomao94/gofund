package v1

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gotomicro/ego-component/eredis"
	"github.com/gotomicro/ego/core/elog"
	"github.com/gotomicro/ego/core/etrace"
	"github.com/hehanpeng/gofund/common/resp"
	"time"
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
	ctx := context.Background()
	// try to obtain my Lock
	lock, err := invoker.RedisStub.LockClient().Obtain(ctx, "my-key", 10*time.Second, eredis.WithLockOptionRetryStrategy(
		eredis.LimitRetry(
			eredis.LinearBackoffRetry(1*time.Second), 6)))
	defer func() {
		lock.Release(ctx)
		invoker.Logger.Info("Release Lock!")
	}()
	if err == eredis.ErrNotObtained {
		invoker.Logger.Error("Could not obtain Lock!", elog.FieldErr(err))
		resp.FailWithMessage("Could not obtain Lock!", c)
		return
	} else if err != nil {
		invoker.Logger.Error("error", elog.FieldErr(err))
		resp.FailWithMessage("error", c)
		return
	}
	invoker.Logger.Info("I have a Lock!")
	//time.Sleep(5 * time.Second)
	resp.Ok(c)
}
