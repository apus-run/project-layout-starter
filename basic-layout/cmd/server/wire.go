//go:build wireinject
// +build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"

	"basic-layout/internal/core/repo"
	"basic-layout/internal/core/repo/dao"
	"basic-layout/internal/core/svc"
	"basic-layout/internal/infra"
	"basic-layout/internal/web"
	"basic-layout/internal/web/handler"
	"basic-layout/pkg/logger"
)

// runApp init server application.
func runApp(log logger.Logger) *gin.Engine {
	panic(wire.Build(
		// 数据库 和 缓存 等
		infra.ProviderSet,

		// DAO 部分
		dao.NewUserDAO,

		// Repository 部分
		repo.NewUserRepository,

		// Service 部分
		svc.NewUserService,

		// Handler 部分
		handler.NewUserHandler,

		// gin 中间件
		web.InitMiddlewares,

		// Web 服务器
		web.InitWebServer,
	))
}
