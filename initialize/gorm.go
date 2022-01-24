package initialize

import (
	"os"

	"github.com/qianxia/blog/model"
	"gorm.io/gorm"
)

func RegisterTables(db *gorm.DB) {
	if err := db.AutoMigrate(
		&model.User{},
	); err != nil {
		os.Exit(0)
	}

}
