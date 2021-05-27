package main

import (
	"github.com/gotomicro/ego"
	"github.com/gotomicro/ego/core/elog"
	"github.com/gotomicro/ego/server/egovernor"
	"waf-srv/invoker"
	"waf-srv/job"
	"waf-srv/router"
)

func main() {
	if err := ego.New().
		Invoker(invoker.Init).
		Job(
			job.InstallComponent(),
		).
		Serve(
			egovernor.Load("server.governor").Build(),
			router.ServeGRPC(),
		).
		Run(); err != nil {
		elog.Panic(err.Error())
	}
}
