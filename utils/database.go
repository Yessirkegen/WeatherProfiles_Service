package utils

import (
	"fmt"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

func GetDBInstance() *gorm.DB {
	once.Do(func() {
		dsn := "host=localhost user=postgres password=yourpassword dbname=userdb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
		var err error
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(fmt.Sprintf("Failed to connect to database: %v", err))
		}
	})
	return db
}
