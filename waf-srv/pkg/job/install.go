package job

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
	"waf-srv/model"
	"waf-srv/pkg/invoker"
	"waf-srv/service"
	"waf-srv/statistics"

	"github.com/gotomicro/ego/core/elog"
	"github.com/gotomicro/ego/core/etrace"
	"github.com/gotomicro/ego/task/ecron"
	"go.uber.org/zap"

	"github.com/gotomicro/ego/task/ejob"
)

var (
	limit = 10000
)

func InstallComponent() *ejob.Component {
	return ejob.DefaultContainer().Build(
		ejob.WithName("install"),
		ejob.WithStartFunc(runInstall),
	)
}

func runInstall(ctx ejob.Context) error {
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

// 超时转发
func CronTtoInfo() ecron.Ecron {
	job := func(ctx context.Context) error {
		// 查找数据 确定协程数
		var ttoInfos []model.TtoInfo
		// 结合系统性能限制拉取条数
		// 如果系统性能差，拉取逻辑需要优化
		err := invoker.Db.Where("tto_status = ? and execute_time <= ?", 0, time.Now()).
			Limit(limit).
			Find(&ttoInfos).Error
		if err != nil {
			invoker.Logger.Error("cronTtoInfo error", zap.Error(err))
			return err
		}
		concurrent := len(ttoInfos)
		if concurrent == 0 {
			invoker.Logger.Info("cronTtoInfo no data")
			return nil
		}
		// 设置接收数据缓存
		ch := make(chan *statistics.RequestResults, 1000)
		var (
			wg          sync.WaitGroup // 发送数据完成
			wgReceiving sync.WaitGroup // 数据处理完成
		)
		wgReceiving.Add(1)
		go statistics.ReceivingResults(uint64(concurrent), ch, &wgReceiving)
		// 执行转发逻辑
		for v, ttoInfo := range ttoInfos {
			wg.Add(1)
			go service.DealTto(uint64(v), ttoInfo, ch, &wg)
		}
		// 等待所有的数据都发送完成
		wg.Wait()
		// 延时1毫秒 确保数据都处理完成了
		time.Sleep(1 * time.Millisecond)
		close(ch)
		// 数据全部处理完成了
		wgReceiving.Wait()
		return nil
	}

	cron := ecron.Load("cron.waf").Build(
		ecron.WithLock(invoker.EcronLocker.NewLock("CronTtoInfo")),
		ecron.WithJob(job))
	return cron
}
