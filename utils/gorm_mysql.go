package utils

import (
	"fmt"
	"net/url"
	"time"

	"github.com/qianxia/blog/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func InitDb(y *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=%s",
		y.MySQL.Username,
		y.MySQL.Password,
		y.MySQL.Host,
		y.MySQL.Port,
		y.MySQL.DbName,
		y.MySQL.Charset,
		url.QueryEscape(y.MySQL.Loc),
	)

	db, _ := gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 150,
	}), &gorm.Config{
		Logger: logger.Default,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "qx_",
			SingularTable: true,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	sqlDB, _ := db.DB()
	// 连接池最大空闲连接数
	sqlDB.SetMaxIdleConns(10)
	//数据库最大连接数
	sqlDB.SetMaxOpenConns(100)
	// 连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	// db.Logger.LogMode(3)
	return db
}
