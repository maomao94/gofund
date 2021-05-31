package main

import (
	"waf-srv/pkg/invoker"
	"waf-srv/pkg/job"
	"waf-srv/pkg/router"

	"github.com/gotomicro/ego"
	"github.com/gotomicro/ego/core/elog"
	"github.com/gotomicro/ego/server/egovernor"
)

//  export EGO_DEBUG=true EGO_NAME=waf-srv && go run main.go --config=config/dev.toml --job=install
func main() {
	if err := ego.New().
		Invoker(invoker.Init).
		//Registry(invoker.EtcdRegistry).
		Job(
			job.InstallComponent(),
		).
		Serve(
			egovernor.Load("server.governor").Build(),
			router.GetRouter(),
		).
		Run(); err != nil {
		elog.Panic(err.Error())
	}
}
