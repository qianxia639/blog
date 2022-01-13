package types

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/qianxia/blog/command"
	"github.com/qianxia/blog/model"
	"github.com/qianxia/blog/utils"
)

type TypeService struct {
	DB *gorm.DB
}

func NewTypeService() TypeService {
	return TypeService{DB: utils.GetDB()}
}

func (ts TypeService) List() ([]model.Type, error) {
	var err error
	Db := utils.GetDB()
	types := make([]model.Type, 4)
	if err = Db.Raw("SELECT id,type_name FROM " + command.DBType).Scan(&types).Error; err != nil {
		return nil, errors.New("查询失败")
	}
	return types, err
}
