package main

import (
	"os"

	"github.com/apus-run/gaia"
	"github.com/apus-run/gaia/transport/grpc"
	"github.com/apus-run/gaia/transport/http"
	"github.com/apus-run/sea-kit/log"

	"gaia-layout/internal/infra"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string

	id, _ = os.Hostname()
)

func newApp(logger log.Logger, gs *grpc.Server, hs *http.Server) *gaia.Gaia {
	return gaia.New(
		gaia.WithID(id),
		gaia.WithName(Name),
		gaia.WithVersion(Version),
		gaia.WithMetadata(map[string]string{}),
		gaia.WithLogger(logger),
		gaia.WithServer(gs, hs),
	)
}

func main() {
	// 初始化配置
	infra.InitConfig()

	// 初始化日志
	logger := log.With(
		log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
	)

	app := runApp(logger)

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
