package server

import (
	"fmt"
	"github.com/Aj002Th/imail/common/config"
	"github.com/Aj002Th/imail/common/logs"
	"github.com/Aj002Th/imail/server/manager"
	"github.com/Aj002Th/imail/server/manager/dal"
	"os"
	"os/signal"
	"syscall"
)

func RunMain(path string) {
	config.Init(path)
	logs.Init()
	dal.Init()
	contentManager := manager.NewContentManager()
	contentManager.Run()

	// 优雅退出
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan)
	waitElegantExit(sigChan)
}

// 优雅退出（退出信号）
// SIGHUP: terminal closed
// SIGINT: Ctrl+C
// SIGTERM: program exit
// SIGQUIT: Ctrl+/
func waitElegantExit(signalChan chan os.Signal) {
	for i := range signalChan {
		switch i {
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			fmt.Println("receive exit signal ", i.String(), ",exit...")
			cleanup()
			os.Exit(0)
		}
	}
}

// 资源清理
func cleanup() {
	logs.Close()
}
