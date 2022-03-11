package initialize

import (
	"os"

	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/model"
	"gorm.io/gorm"
)

func RegisterTables(db *gorm.DB) {
	if err := db.AutoMigrate(
		&model.User{},
		&model.Type{},
		&model.Tag{},
		&model.Blog{},
		&model.Comment{},
	); err != nil {
		global.QX_LOG.Fatalf("表自动迁移失败,err: %s", err)
		os.Exit(0)
	}
}

func InitData() {
	global.QX_DB.Create(&model.Type{Id: 1, TypeName: "Golang"})
	global.QX_DB.Create(&model.Type{Id: 2, TypeName: "日志"})
	global.QX_DB.Create(&model.Type{Id: 3, TypeName: "数据库"})
	global.QX_DB.Create(&model.Type{Id: 4, TypeName: "前端"})
}
