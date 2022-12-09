package db

import (
	"backend/common"
	"sync"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Type        = "mysql"
	Url         = "username:password@tcp(127.0.0.1:3306)/beekeeper?charset=utf8mb4&loc=Local"
	MaxIdleConn = 2
	MaxOpenConn = 10
	MaxLifetime = time.Hour
)

var db *gorm.DB
var dbOnce = sync.Once{}

func GetDB() *gorm.DB {
	dbOnce.Do(func() {
		if Type == "mysql" {
			d, err := gorm.Open(mysql.Open(Url), &gorm.Config{})
			common.PanicNotNull(err)
			sqlDB, err := d.DB()
			common.PanicNotNull(err)
			sqlDB.SetMaxIdleConns(MaxIdleConn)
			sqlDB.SetMaxOpenConns(MaxOpenConn)
			sqlDB.SetConnMaxLifetime(MaxLifetime)
			db = d
		}
	})

	return db
}
