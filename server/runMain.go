package server

import (
	"github.com/Aj002Th/imail/common/config"
	"github.com/Aj002Th/imail/common/logs"
	"github.com/Aj002Th/imail/server/manager"
	"github.com/Aj002Th/imail/server/manager/dal"
)

func RunMain(path string) {
	config.Init(path)
	logs.Init()
	dal.Init()
	contentManager := manager.NewContentManager()
	contentManager.Run()
}
