package cmd

import (
	"github.com/Aj002Th/imail/common/config"
	"github.com/Aj002Th/imail/common/logs"
	"github.com/Aj002Th/imail/server/manager"
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
	if len(args) == 0 {
		slog.Error("args is empty")
		return
	}
	function, ok := cmdMap[args[0]]
	if !ok {
		slog.Error("function not found")
		return
	}

	config.Init(ConfigPath)
	logs.Init()
	dal.Init()
	function(args[1:])
}

var cmdMap = map[string]func([]string){
	"clear": clearCmd,
	"send":  sendCmd,
}

func clearCmd(args []string) {
	slog.Info("run cmd: imail content clearCmd")
	db := dal.GetDB()
	err := db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.Content{}).Error
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}
}

func sendCmd(args []string) {
	slog.Info("run cmd: imail content sendCmd")
	contentManager := manager.NewContentManager()
	contentManager.CatchAndSend()
}
