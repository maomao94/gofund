module waf-srv

go 1.16

require (
	github.com/gin-gonic/gin v1.7.2
	github.com/gotomicro/ego v0.5.6
	github.com/gotomicro/ego-component/egorm v0.2.0
	github.com/gotomicro/ego-component/eredis v0.2.2
	github.com/hehanpeng/gofund/common v0.0.0-00010101000000-000000000000
	github.com/hehanpeng/gofund/proto v0.0.0-00010101000000-000000000000
	go.uber.org/zap v1.17.0
	google.golang.org/grpc v1.29.1
)

replace (
	github.com/hehanpeng/gofund/common => ../common
	github.com/hehanpeng/gofund/proto => ../proto
)
