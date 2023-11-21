package cmd

import (
	"fmt"
	"github.com/Aj002Th/imail/common/config"
	"github.com/Aj002Th/imail/common/logs"
	"github.com/Aj002Th/imail/server/manager/dal"
	"github.com/Aj002Th/imail/server/manager/dal/model"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
	"log/slog"
)

func init() {
	rootCmd.AddCommand(contentCmd)
}

var contentCmd = &cobra.Command{
	Use: "content",
	Run: contentHandle,
}

func contentHandle(cmd *cobra.Command, args []string) {
	fmt.Println("imail content db")
	if len(args) == 0 {
		slog.Error("args is empty")
		return
	}
	function, ok := cmdMap[args[0]]
	if !ok {
		slog.Error("function not found")
		return
	}
	function()
}

var cmdMap = map[string]func(){
	"clear": clearCmd,
}

func clearCmd() {
	fmt.Println("imail content clearCmd")
	config.Init(ConfigPath)
	logs.Init()
	dal.Init()
	db := dal.GetDB()
	err := db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.Content{}).Error
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}
}
