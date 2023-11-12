package server

import (
	"time"

	"github.com/apus-run/gaia/middleware/recovery"
	"github.com/apus-run/gaia/transport/grpc"
	"github.com/spf13/viper"

	hv1 "gaia-layout/api/helloworld/v1"
	uv1 "gaia-layout/api/user/v1"
	"gaia-layout/internal/rpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(greeter *rpc.GreeterService, user *rpc.UserService) *grpc.Server {
	// 在初始化模块的时候再读配置信息
	type Config struct {
		Network string `yaml:"network,omitempty"`
		Addr    string `yaml:"addr,omitempty"`
		Timeout string `yaml:"timeout,omitempty"`
	}
	var cfg Config
	err := viper.UnmarshalKey("server.grpc", &cfg)
	if err != nil {
		panic(err)
	}

	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
		),
	}
	if cfg.Network != "" {
		opts = append(opts, grpc.Network(cfg.Network))
	}
	if cfg.Addr != "" {
		opts = append(opts, grpc.Address(cfg.Addr))
	}
	if cfg.Timeout != "" {
		timeout, err := time.ParseDuration(cfg.Timeout)
		if err == nil {
			opts = append(opts, grpc.Timeout(timeout))
		}
	}
	srv := grpc.NewServer(opts...)
	hv1.RegisterGreeterServer(srv, greeter)
	uv1.RegisterUserServer(srv, user)
	return srv
}
