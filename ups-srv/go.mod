module ups-srv

go 1.16

require (
	github.com/HdrHistogram/hdrhistogram-go v1.1.0 // indirect
	github.com/envoyproxy/go-control-plane v0.9.9-0.20210512163311-63b5d3c536b0 // indirect
	github.com/gin-gonic/gin v1.7.2
	github.com/go-pay/gopay v1.5.49
	github.com/go-resty/resty/v2 v2.6.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/gotomicro/ego v0.5.8
	github.com/gotomicro/ego-component/eetcd v0.2.1
	github.com/gotomicro/ego-component/egorm v0.2.1
	github.com/gotomicro/ego-component/eredis v0.2.2
	github.com/hehanpeng/gofund/common v0.0.0-20210604012407-335b57ce0918
	github.com/hehanpeng/gofund/proto v0.0.0-20210604012407-335b57ce0918
	github.com/prometheus/client_golang v1.11.0 // indirect
	github.com/tklauser/go-sysconf v0.3.6 // indirect
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	go.uber.org/zap v1.17.0
	golang.org/x/lint v0.0.0-20210508222113-6edffad5e616 // indirect
	golang.org/x/net v0.0.0-20210716203947-853a461950ff // indirect
	golang.org/x/sys v0.0.0-20210630005230-0f9fa26af87c // indirect
	golang.org/x/tools v0.1.5 // indirect
	google.golang.org/grpc v1.38.1
)

replace (
	github.com/hehanpeng/gofund/common => ../common
	github.com/hehanpeng/gofund/proto => ../proto
	waf-srv => ../waf-srv
)
