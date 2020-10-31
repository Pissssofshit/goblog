package database

import (
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var once sync.Once

func GetInstance() *gorm.DB {
	once.Do(func() {
		dsn := "root:root@tcp(127.0.0.1:8889)/blog_service?charset=utf8mb4&parseTime=True&loc=Local"
		db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	})
	return db
}
