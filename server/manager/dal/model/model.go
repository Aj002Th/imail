package model

import (
	"github.com/Aj002Th/imail/server/catcher"
	"gorm.io/gorm"
	"time"
)

type Content struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index;index:idx,unique"`

	catcher.Content

	Sended bool `gorm:"default:false"`
}
