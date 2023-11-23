package model

import (
	"github.com/Aj002Th/imail/server/catcher"
	"gorm.io/gorm"
)

type Content struct {
	gorm.Model
	catcher.Content

	Sended bool `gorm:"default:false"`
}
