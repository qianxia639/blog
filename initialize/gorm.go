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
		&model.Role{},
	); err != nil {
		global.QX_LOG.Fatalf("表自动迁移失败,err: %s", err)
		os.Exit(0)
	}
}
