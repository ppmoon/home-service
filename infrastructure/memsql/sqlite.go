package memsql

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"sync"
)

const sqliteMemoryDsn = "file::memory:?cache=shared"

var db *gorm.DB
var memsqlOnce sync.Once

func GetDB() *gorm.DB {
	memsqlOnce.Do(func() {
		var err error
		db, err = gorm.Open(sqlite.Open(sqliteMemoryDsn), &gorm.Config{
			SkipDefaultTransaction: true,
		})
		if err != nil {
			panic(err)
		}
	})
	return db
}
