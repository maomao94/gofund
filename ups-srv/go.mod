module ups-srv

go 1.16

require (
	github.com/gin-gonic/gin v1.7.2
	github.com/go-pay/gopay v1.5.49
	github.com/gotomicro/ego v0.5.7
	github.com/gotomicro/ego-component/egorm v0.2.1
	github.com/gotomicro/ego-component/eredis v0.2.2
	github.com/hehanpeng/gofund/common v0.0.0-20210604012407-335b57ce0918
	github.com/hehanpeng/gofund/proto v0.0.0-20210604012407-335b57ce0918
	go.uber.org/zap v1.17.0
	google.golang.org/grpc v1.38.0
	waf-srv v0.0.0-00010101000000-000000000000
)

replace (
	github.com/hehanpeng/gofund/common => ../common
	github.com/hehanpeng/gofund/proto => ../proto
	waf-srv => ../waf-srv
)
