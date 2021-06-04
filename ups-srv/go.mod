module ups-srv

go 1.16

require (
	github.com/gin-gonic/gin v1.7.2
	github.com/gotomicro/ego v0.5.7
	github.com/gotomicro/ego-component/egorm v0.2.1
	github.com/gotomicro/ego-component/eredis v0.2.2
	github.com/iGoogle-ink/gopay v1.5.40
	go.uber.org/zap v1.17.0
	google.golang.org/grpc v1.38.0
)

replace (
	github.com/hehanpeng/gofund/common => ../common
	github.com/hehanpeng/gofund/proto => ../proto
	waf-srv => ../waf-srv
)
