package config

import "github.com/spf13/viper"

func GetCronTab() string {
	return viper.GetString("manager.crontab")
}

func IsImmediate() bool {
	return viper.GetBool("manager.immediate")
}

func IsIgnoreEmptyMessage() bool {
	return viper.GetBool("manager.ignoreEmptyMessage")
}
