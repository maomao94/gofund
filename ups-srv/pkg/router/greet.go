package router

import (
	"context"
	"errors"
	"log"
	"ups-srv/pkg/invoker"

	"github.com/hehanpeng/gofund/proto/fund/gen/upssrv"

	"github.com/hehanpeng/gofund/proto/fund/gen/errcodepb"
)

func (s *Ups) SayHello(ctx context.Context, in *upssrv.HelloRequest) (*upssrv.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	//return &helloworld.HelloReply{Message: "Hello " + in.GetName()}, nil
	return nil, invoker.Error(errcodepb.ErrCode_OK, errors.New(errcodepb.ErrCode_OK.String()))
}
