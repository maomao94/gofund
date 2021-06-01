package v1

import (
	"context"
	"time"
	"waf-srv/model"
	"waf-srv/pkg/invoker"

	"github.com/hehanpeng/gofund/common/global/api"

	"github.com/gin-gonic/gin"
	"github.com/gotomicro/ego-component/eredis"
	"github.com/gotomicro/ego/core/elog"
	"github.com/gotomicro/ego/core/etrace"
)

func TestUps(c *gin.Context) {
	span, ctx := etrace.StartSpanFromContext(context.Background(), "callHTTP()")
	defer span.Finish()

	req := invoker.UpsHttpComp.R()
	// Inject traceId Into Header
	c1 := etrace.HeaderInjector(ctx, req.Header)

	info, err := req.SetContext(c1).Get("/api/hello")
	if err != nil {
		api.Fail(c)
		return
	}
	api.OkWithData(info.String(), c)
}

func HelloLock(c *gin.Context) {
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
		api.FailWithMessage("Could not obtain Lock!", c)
		return
	} else if err != nil {
		invoker.Logger.Error("error", elog.FieldErr(err))
		api.FailWithMessage("error", c)
		return
	}
	invoker.Logger.Info("I have a Lock!")
	//time.Sleep(5 * time.Second)
	api.Ok(c)
}

func Hello(c *gin.Context) {
	var ttoInfo model.TtoInfo
	_ = c.ShouldBindJSON(&ttoInfo)
	invoker.Logger.Infof("hello tto: %v", ttoInfo)
	time.Sleep(3 * time.Second)
	api.Fail(c)
}
