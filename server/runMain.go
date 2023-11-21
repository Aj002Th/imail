package server

import (
	"github.com/Aj002Th/imail/common/config"
	"github.com/Aj002Th/imail/server/manager"
	"github.com/Aj002Th/imail/server/manager/dal"
)

func RunMain(path string) {
	config.Init(path)
	dal.Init()
	manager := manager.NewContentManager()
	manager.Run()
}
