package config

import (
	"github.com/spf13/viper"
)

// IsDebug 判断是不是debug环境
func IsDebug() bool {
	env := viper.GetString("global.env")
	return env == "debug"
}