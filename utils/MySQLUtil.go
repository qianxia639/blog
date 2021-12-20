package utils

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/qianxia/blog/model"
)

var Db *gorm.DB

func InitDb(y *model.Config) *gorm.DB {

	driverName := y.MySQL.DriverName
	args := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		y.MySQL.Username,
		y.MySQL.Password,
		y.MySQL.Host,
		y.MySQL.Port,
		y.MySQL.DbName,
		y.MySQL.Charset)

	db, _ := gorm.Open(driverName, args)
	// 设置最大空闲连接数
	db.DB().SetMaxIdleConns(10)
	// 设置最大连接数
	db.DB().SetMaxOpenConns(30)
	// 设置可重用连接的最长时间
	db.DB().SetConnMaxLifetime(time.Minute * 30)

	db.LogMode(true)
	Db = db
	return Db
}

func GetDB() *gorm.DB {
	return Db
}
