package infra

import (
	"basic-layout/pkg/logger"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// InitLogger 日志 和 日志文件
func InitLogger() logger.Logger {
	// 小技巧, 直接使用 pkg logger 本身的配置结构体来序列化
	cfg := logger.DefaultOptions()
	err := viper.UnmarshalKey("log", &cfg)
	if err != nil {
		panic(err)
	}

	return logger.NewLogger(
		logger.WithMode(cfg.Mode),
		logger.WithLogLevel(cfg.LogLevel),
		logger.WithEncoding(cfg.Encoding),

		logger.WithFilename(cfg.LogFilename),
		logger.WithMaxSize(cfg.MaxSize),
		logger.WithMaxBackups(cfg.MaxBackups),
		logger.WithMaxAge(cfg.MaxAge),
		logger.WithCompress(cfg.Compress),
	)
}

// InitZapLogger 日志
func InitZapLogger() logger.Logger {
	// 小技巧, 直接使用 zap 本身的配置结构体来序列化
	cfg := zap.NewDevelopmentConfig()
	err := viper.UnmarshalKey("log", &cfg)
	if err != nil {
		panic(err)
	}
	l, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	return logger.NewZapLogger(l)
}
