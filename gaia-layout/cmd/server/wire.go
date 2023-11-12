//go:build wireinject
// +build wireinject

package main

import (
	"github.com/apus-run/gaia"
	"github.com/apus-run/sea-kit/log"
	"github.com/google/wire"

	"gaia-layout/internal/core/repo"
	"gaia-layout/internal/core/repo/dao"
	"gaia-layout/internal/core/svc"
	"gaia-layout/internal/infra"
	"gaia-layout/internal/rpc"
	"gaia-layout/internal/server"
)

// runApp init server application.
func runApp(log.Logger) *gaia.Gaia {
	panic(wire.Build(
		// 数据库 和 缓存 等
		infra.ProviderSet,

		// DAO 部分
		dao.NewUserDAO,

		// Repository 部分
		repo.NewUserRepository,

		// Service 部分
		svc.NewUserService,

		rpc.NewGreeterService,
		rpc.NewUserService,
		server.ProviderSet,
		newApp,
	))
}
