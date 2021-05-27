package main

import (
	"waf-srv/invoker"
	"waf-srv/job"
	"waf-srv/router"

	"github.com/gotomicro/ego"
	"github.com/gotomicro/ego/core/elog"
	"github.com/gotomicro/ego/server/egovernor"
)

//  export EGO_DEBUG=true && go run main.go --config=config/dev.toml
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
