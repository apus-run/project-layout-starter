package infra

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"basic-layout/internal/core/repo/dao/model"
)

func InitDB() *gorm.DB {
	// 在初始化模块的时候再读配置信息
	type Config struct {
		DSN string `yaml:"dsn"`
	}
	var cfg Config
	err := viper.UnmarshalKey("db", &cfg)
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(
		mysql.Open(cfg.DSN),
		&gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
		},
	)
	if err != nil {
		panic(err)
	}

	// 没有开启debug模式，不打印sql
	db.Logger = logger.Default.LogMode(logger.Silent)

	// 为了方便，我们这里直接把表初始化放在这里
	model.InitTables(db)

	return db
}
