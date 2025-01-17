package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"la-blog-go/global"
	"la-blog-go/model"
)

func ConnectToSQLiteDB(dbPath string, autoMigrate bool) {
	dsn := dbPath
	db, _ := gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	global.DB = db
	if autoMigrate {
		err := db.AutoMigrate(&model.Article{}, &model.Category{}, &model.Tag{})
		if err != nil {
			return
		}
	}
}
