module waf-srv

go 1.16

require (
	github.com/gin-gonic/gin v1.7.2
	github.com/gotomicro/ego v0.6.2
	github.com/gotomicro/ego-component/egorm v0.2.1
	github.com/gotomicro/ego-component/eredis v0.2.2
	github.com/hehanpeng/gofund/common v0.0.0-20210604012407-335b57ce0918
	github.com/hehanpeng/gofund/proto v0.0.0-20210604012407-335b57ce0918
	go.uber.org/zap v1.17.0
	golang.org/x/text v0.3.6
	google.golang.org/grpc v1.39.0
)

replace (
	github.com/hehanpeng/gofund/common => ../common
	github.com/hehanpeng/gofund/proto => ../proto
)
