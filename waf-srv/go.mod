module waf-srv

go 1.16

require (
	github.com/gin-gonic/gin v1.7.1
	github.com/google/uuid v1.1.2 // indirect
	github.com/gotomicro/ego v0.5.6
	github.com/gotomicro/ego-component/eetcd v0.2.1
	github.com/gotomicro/ego-component/egorm v0.2.0
	github.com/gotomicro/ego-component/eredis v0.2.2
	github.com/hehanpeng/gofund/proto v0.0.0-00010101000000-000000000000
	go.uber.org/zap v1.15.0
	golang.org/x/lint v0.0.0-20201208152925-83fdc39ff7b5 // indirect
	golang.org/x/tools v0.1.0 // indirect
	google.golang.org/grpc v1.29.1
)

replace github.com/hehanpeng/gofund/proto => ../proto
