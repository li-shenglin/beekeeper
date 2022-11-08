package db

import (
	"backend/common"
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

func GetDB() *gorm.DB {
	if Type == "mysql" {
		db, err := gorm.Open(mysql.Open(Url), &gorm.Config{})
		common.PanicNotNull(err)
		sqlDB, err := db.DB()
		common.PanicNotNull(err)
		sqlDB.SetMaxIdleConns(MaxIdleConn)
		sqlDB.SetMaxOpenConns(MaxOpenConn)
		sqlDB.SetConnMaxLifetime(MaxLifetime)
		return db
	}
	return nil
}
