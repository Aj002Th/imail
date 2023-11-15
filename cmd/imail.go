package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

var (
	ConfigPath string
)

func init() {
	cobra.OnInitialize(initConfig)

	// 命令行参数
	// 配置文件路径
	rootCmd.PersistentFlags().StringVar(
		&ConfigPath,
		"config",
		"./imail.yaml",
		"config file (default is ./imail.yaml)")
}

func initConfig() {
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Panicln(err)
	}
}

var rootCmd = &cobra.Command{
	Use: "imail",
	Run: imail,
}

func imail(cmd *cobra.Command, args []string) {
	fmt.Println("imail rootCmd")
}
