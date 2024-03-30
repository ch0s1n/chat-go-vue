package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"sync"
)

var once sync.Once
var MyDb *gorm.DB

func InitDB() *gorm.DB {
	once.Do(func() {
		dsn := "root:123456@tcp(127.0.0.1:3306)/chat?charset=utf8mb4&parseTime=True&loc=Local"
		mydb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
		if err != nil {
			log.Fatal("failed to connect database")
		}
		MyDb = mydb
	})
	return MyDb
}
