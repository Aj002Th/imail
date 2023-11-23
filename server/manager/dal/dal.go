package dal

import (
	"github.com/Aj002Th/imail/common/logs"
	"github.com/Aj002Th/imail/server/manager/dal/model"
	"github.com/Aj002Th/imail/server/manager/dal/query"
	slogGorm "github.com/orandin/slog-gorm"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log/slog"
)

var db *gorm.DB

func Init() {
	var err error
	gormLogger := slogGorm.New(
		slogGorm.WithLogger(logs.Logger), // Optional, use slog.Default() by default
		slogGorm.WithTraceAll(),          // trace all messages
	)

	db, err = gorm.Open(sqlite.Open("./data/sqlite/manager.db"), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}

	err = db.AutoMigrate(&model.Content{})
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}

	query.SetDefault(db)
}

func GetDB() *gorm.DB {
	return db
}
