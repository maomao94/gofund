package invoker

import (
	"fmt"

	"github.com/gotomicro/ego/client/egrpc/resolver"

	"github.com/gotomicro/ego/client/ehttp"

	"github.com/gotomicro/ego/server/egin"

	"github.com/gotomicro/ego-component/eetcd"
	"github.com/gotomicro/ego-component/eetcd/registry"
	"github.com/gotomicro/ego-component/egorm"
	"github.com/gotomicro/ego-component/eredis"
	"github.com/gotomicro/ego/core/elog"
	"github.com/hehanpeng/gofund/proto/fund/gen/errcodepb"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	Logger       *elog.Component
	Gin          *egin.Component
	Db           *egorm.Component
	RedisStub    *eredis.Component
	WafHttpComp  *ehttp.Component
	EtcdClient   *eetcd.Component
	EtcdRegistry *registry.Component
)

func Init() error {
	Logger = elog.DefaultLogger
	Gin = egin.Load("server.http").Build()
	Db = egorm.Load("mysql.ups").Build()
	RedisStub = eredis.Load("redis.ups").Build(eredis.WithStub())
	WafHttpComp = ehttp.Load("http.waf").Build()
	EtcdClient = eetcd.Load("etcd").Build()
	EtcdRegistry = registry.Load("registry").Build(registry.WithClientEtcd(EtcdClient))

	// 必须注册在grpc前面
	resolver.Register("etcd", EtcdRegistry)
	return nil
}

// 记录grpc error信息
func Error(code errcodepb.ErrCode, err error) error {
	Logger.Error("grpc error", zap.Int32("code", int32(code)), zap.Error(err))
	var cause string
	if err != nil {
		cause = err.Error()
	}
	return status.Error(codes.Code(code), fmt.Sprintf("error name: %s, cause: %s", errcodepb.ErrCode_name[int32(code)], cause))
}
