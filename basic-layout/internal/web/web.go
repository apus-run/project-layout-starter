package web

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"basic-layout/internal/web/handler"
	"basic-layout/internal/web/middleware/accesslog"
	"basic-layout/pkg/logger"
)

// InitWebServer gin 服务
func InitWebServer(funcs gin.HandlersChain, userHandler handler.UserHandler) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	server := gin.Default()
	server.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "pong",
		})
	})
	server.Use(funcs...)
	userHandler.RegisterRoutes(server)
	return server
}

// InitMiddlewares gin中间件
func InitMiddlewares() gin.HandlersChain {
	return []gin.HandlerFunc{
		accesslog.NewBuilder(
			func(ctx context.Context, al *accesslog.AccessLog) {
				logger.L().Errorf("Gin 收到请求", logger.Any("req", al))
			}).
			AllowReqBody().
			AllowRespBody().
			IgnoreRoutes("/login").
			Build(),
	}
}
