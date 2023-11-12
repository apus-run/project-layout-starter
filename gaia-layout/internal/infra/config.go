package infra

import (
	"fmt"

	"github.com/apus-run/sea-kit/log"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func InitConfig() {
	flagconf := pflag.String("conf", "config/dev.yaml", "config path, eg: -conf config.yaml")
	pflag.Parse()
	// 直接指定文件路径
	viper.SetConfigFile(*flagconf)
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println(in.Name, in.Op)
	})
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Infof("Using conf file: %s [%s]\n", viper.ConfigFileUsed(), err)
		}
		panic(err)
	}
}
