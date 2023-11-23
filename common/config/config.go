package config

import (
	"github.com/spf13/viper"
)

// Init 读配置文件, 获取配置信息
func Init(path string) {
	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
