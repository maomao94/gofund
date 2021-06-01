package job

import (
	"context"
	"errors"
	"fmt"
	"github.com/gotomicro/ego/core/elog"
	"github.com/gotomicro/ego/core/etrace"
	"github.com/gotomicro/ego/task/ecron"
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
		return nil
	}

	cron := ecron.Load("cron.waf").Build(
		ecron.WithLock(invoker.EcronLocker.NewLock("CronJob2")),
		ecron.WithJob(job))
	return cron
}
