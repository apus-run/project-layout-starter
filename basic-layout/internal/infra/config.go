package infra

import (
	"fmt"
	"log"
	"time"

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
			log.Printf("Using conf file: %s [%s]\n", viper.ConfigFileUsed(), err)
		}
		panic(err)
	}
}

func InitConfigRemote() {
	err := viper.AddRemoteProvider("etcd3", "http://127.0.0.1:12379", "/server")
	if err != nil {
		panic(err)
	}
	viper.SetConfigType("yaml")
	err = viper.ReadRemoteConfig()
	if err != nil {
		panic(err)
	}
	go func() {
		for {
			err = viper.WatchRemoteConfig()
			if err != nil {
				log.Println(err)
				return
			}
			// 睡个一秒钟
			time.Sleep(time.Second)
		}
	}()
}
