package config

import "github.com/spf13/viper"

type EmailSender struct {
	Nickname string
	Host     string
	Port     int
	Username string
	Password string
}

func GetEmailSenderForMessager() EmailSender {
	return EmailSender{
		Nickname: viper.GetString("messager.email.sender.nickname"),
		Host:     viper.GetString("messager.email.sender.host"),
		Port:     viper.GetInt("messager.email.sender.port"),
		Username: viper.GetString("messager.email.sender.username"),
		Password: viper.GetString("messager.email.sender.password"),
	}
}

func GetEmailReceiversForMessager() []string {
	return viper.GetStringSlice("messager.email.receiver.users")
}
