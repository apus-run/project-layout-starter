package infra

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

func InitRDB() redis.Cmdable {
	// 在初始化模块的时候再读配置信息
	type Config struct {
		Addr         string `yaml:"addr,omitempty"`
		Password     string `yaml:"password,omitempty"`
		DB           int    `yaml:"db,omitempty"`
		ReadTimeout  string `yaml:"read_timeout,omitempty"`
		WriteTimeout string `yaml:"write_timeout,omitempty"`
	}

	var cfg Config
	err := viper.UnmarshalKey("data.redis", &cfg)
	if err != nil {
		panic(err)
	}
	readTimeout, err := time.ParseDuration(cfg.ReadTimeout)
	if err != err {
		panic(err)
	}
	writeTimeout, err := time.ParseDuration(cfg.WriteTimeout)
	if err != err {
		panic(err)
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,

		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}

	return rdb
}
