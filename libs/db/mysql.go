package db

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
)

var DB *gorm.DB

func InitDB(env string, write string, read string) {
	// https://gorm.io/zh_CN/docs/dbresolver.html
	// 建立读写分离连接池
	dsn_master := write
	dsn_slave := read

	// log level
	var level logger.LogLevel
	if env == "local" || env == "dev" || env == "test" {
		level = logger.Info
	} else {
		level = logger.Silent
	}
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second / 4, // Slow SQL threshold
			LogLevel:      level,           // Log level
		},
	)
	DB, _ = gorm.Open(mysql.Open(dsn_master), &gorm.Config{Logger: newLogger})
	DB.Use(dbresolver.Register(dbresolver.Config{
		Sources:  []gorm.Dialector{mysql.Open(dsn_master)},
		Replicas: []gorm.Dialector{mysql.Open(dsn_slave)},
		Policy:   dbresolver.RandomPolicy{},
	}).
		SetConnMaxIdleTime(time.Hour).
		SetConnMaxLifetime(24 * time.Hour).
		SetMaxIdleConns(100).
		SetMaxOpenConns(200))
}
