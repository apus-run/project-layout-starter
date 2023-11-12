package rpc

import (
	"context"

	v1 "gaia-layout/api/helloworld/v1"
)

var _ v1.GreeterServer = (*GreeterService)(nil)

type GreeterService struct {
	v1.UnimplementedGreeterServer
}

func NewGreeterService() *GreeterService {
	return &GreeterService{}
}

// SayHello implements helloworld.GreeterServer.
func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	return &v1.HelloReply{Message: "Hello " + in.Name}, nil
}
