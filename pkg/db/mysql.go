package db

import (
	"fmt"
	"1536509937/ku-bbs/pkg/config"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormDefaultLogger "gorm.io/gorm/logger"
)

var DB *gorm.DB

func init() {
	c := config.Conf.DB
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.Name, c.Pass, c.Host, c.Port, c.DB)
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         255,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		Logger:                 gormDefaultLogger.Default.LogMode(gormDefaultLogger.Info),
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Fatal(err)
	}
	db, err := gormDB.DB()
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(40)
	db.SetConnMaxLifetime(time.Hour)
	DB = gormDB
}
