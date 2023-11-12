package main

import (
	"net/http"
	"os"
	"time"

	_ "go.uber.org/automaxprocs"

	"basic-layout/internal/infra"
	"basic-layout/pkg/logger"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string

	id, _ = os.Hostname()
)

func main() {
	// 初始化配置文件
	infra.InitConfig()

	// 初始化日志
	log := infra.InitZapLogger()
	log.With(
		logger.String("server.id", id),
		logger.String("server.name", Name),
		logger.String("server.version", Version),
	)
	defer log.Sync()

	// 启动Web服务
	router := runApp(log)
	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: router,

		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Infof("server listen at %s", server.Addr)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("服务启动失败: %v", err)
	}
}
