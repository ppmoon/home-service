package sqlite

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const dbName = "home-service.db"

var DB *gorm.DB

func InitGorm() {
	var err error
	DB, err = gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}
