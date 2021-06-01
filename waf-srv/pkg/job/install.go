package job

import (
	"context"
	"errors"
	"fmt"
	"github.com/gotomicro/ego/core/elog"
	"github.com/gotomicro/ego/core/etrace"
	"github.com/gotomicro/ego/task/ecron"
	"go.uber.org/zap"
	"time"
	"waf-srv/model"
	"waf-srv/pkg/invoker"

	"github.com/gotomicro/ego/task/ejob"
)

func InstallComponent() *ejob.Component {
	return ejob.DefaultContainer().Build(
		ejob.WithName("install"),
		ejob.WithStartFunc(runInstall),
	)
}

func runInstall(ctx context.Context) error {
	models := []interface{}{
		//&model.SysApi{},
		&model.TtoInfo{},
	}
	gormdb := invoker.Db.Debug()
	err := gormdb.AutoMigrate(models...)
	if err != nil {
		return err
	}
	fmt.Println("create table ok")
	return nil
}

// 异常任务
func CronJob1() ecron.Ecron {
	job := func(ctx context.Context) error {
		elog.Info("info job1", elog.FieldTid(etrace.ExtractTraceID(ctx)))
		elog.Warn("warn job1", elog.FieldTid(etrace.ExtractTraceID(ctx)))
		fmt.Println("run job1", elog.FieldTid(etrace.ExtractTraceID(ctx)))
		return errors.New("exec job1 error")
	}

	cron := ecron.Load("cron.waf").Build(
		ecron.WithLock(invoker.EcronLocker.NewLock("CronJob1")),
		ecron.WithJob(job))
	return cron
}

// 正常任务
func CronJob2() ecron.Ecron {
	job := func(ctx context.Context) error {
		elog.Info("info job2", elog.FieldTid(etrace.ExtractTraceID(ctx)))
		elog.Warn("warn job2", elog.FieldTid(etrace.ExtractTraceID(ctx)))
		fmt.Println("run job2", elog.FieldTid(etrace.ExtractTraceID(ctx)))
		req := invoker.CallSrvHttpComp["wafs"].R()
		// Inject traceId Into Header
		c1 := etrace.HeaderInjector(ctx, req.Header)
		info, err := req.SetContext(c1).Get("/api/hello")
		if err != nil {
			return err
		}
		elog.Info(info.String())
		return nil
	}

	cron := ecron.Load("cron.waf").Build(
		ecron.WithLock(invoker.EcronLocker.NewLock("CronJob2")),
		ecron.WithJob(job))
	return cron
}

// 超时转发
func CronTtoInfo() ecron.Ecron {
	job := func(ctx context.Context) error {
		statTime := uint64(time.Now().UnixNano())
		var ttoinfos []model.TtoInfo
		//   example.createCriteria()
		//                    .andEqualTo(TtoInf.TTO_STATUS, ForwardParamConstant.TTOINF_STATUS_UNDEAL)
		//                    .andLessThanOrEqualTo(TtoInf.EXCUTE_TIME, new Date())

		err := invoker.Db.Where("tto_status = ? and execute_time <= ?", 0, time.Now()).Find(&ttoinfos).Error
		time.Sleep(10 * time.Second)
		if err != nil {
			invoker.Logger.Error("CronTtoInfo error: ", zap.Error(err))
		}
		endTime := uint64(time.Now().UnixNano())
		requestTimeFloat := float64(endTime-statTime) / 1e9
		invoker.Logger.Infof("CronTtoInfo 耗时: %4.0fs", requestTimeFloat)
		return nil
	}

	cron := ecron.Load("cron.waf").Build(
		ecron.WithLock(invoker.EcronLocker.NewLock("CronTtoInfo")),
		ecron.WithJob(job))
	return cron
}

// header 打印表头信息
func header() {
	// 打印的时长都为毫秒 总请数
	invoker.Logger.Info("─────┬───────┬───────┬───────┬────────┬────────┬────────┬────────┬────────┬────────┬────────")
	invoker.Logger.Info(" 耗时│ 并发数│ 成功数│ 失败数│   qps  │最长耗时│最短耗时│平均耗时│下载字节│字节每秒│ 错误码")
	invoker.Logger.Info("─────┼───────┼───────┼───────┼────────┼────────┼────────┼────────┼────────┼────────┼────────")
	return
}
