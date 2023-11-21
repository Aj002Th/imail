package dal

import (
	"github.com/Aj002Th/imail/server/manager/dal/model"
	"github.com/Aj002Th/imail/server/manager/dal/query"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	var err error
	db, err = gorm.Open(sqlite.Open("./data/sqlite/manager.db"), &gorm.Config{})
	db.AutoMigrate(&model.Content{})
	if err != nil {
		panic(err)
	}
	query.SetDefault(db)
}
