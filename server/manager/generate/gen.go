package main

import (
	"github.com/Aj002Th/imail/server/manager/dal/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./server/manager/dal/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	db, err := gorm.Open(sqlite.Open("./data/sqlite/manager.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	g.UseDB(db) // reuse your gorm db

	// Generate basic type-safe DAO API for struct `model.User` following conventions
	g.ApplyBasic(model.Content{})

	// Generate Type Safe API with Dynamic SQL defined on Querier interface for model
	g.ApplyInterface(func(model.Querier) {}, model.Content{})

	// Generate the code
	g.Execute()
}
