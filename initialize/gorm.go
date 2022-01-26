package initialize

import (
	"os"

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
		os.Exit(0)
	}

}
