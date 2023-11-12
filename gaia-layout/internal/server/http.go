package server

import (
	"net/http"

	"github.com/apus-run/gaia/middleware/recovery"
	"github.com/apus-run/gaia/pkg/ginx"
	httpServer "github.com/apus-run/gaia/transport/http"
	"github.com/gin-gonic/gin"

	hv1 "gaia-layout/api/helloworld/v1"
	uv1 "gaia-layout/api/user/v1"
	"gaia-layout/internal/rpc"
)

func NewRouter() *gin.Engine {
	g := gin.New()
	// 使用gaia中间件
	g.Use(ginx.Middlewares(recovery.Recovery()))

	g.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(http.StatusOK, map[string]string{"welcome": name})
	})

	return g
}

func NewHTTPServer(greeter *rpc.GreeterService, user *rpc.UserService) *httpServer.Server {
	gin.SetMode("release")
	router := NewRouter()

	// http server
	srv := httpServer.NewServer(
		httpServer.Address(":8000"),
		httpServer.Middleware(
			recovery.Recovery(),
		),
	)
	srv.Handler = router

	hv1.RegisterGreeterHTTPServer(router, greeter)
	uv1.RegisterUserHTTPServer(router, user)

	return srv
}
