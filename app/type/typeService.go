package app

import (
	"errors"

	"github.com/qianxia/blog/command"
	"github.com/qianxia/blog/model"
	"github.com/qianxia/blog/utils"
)

type TypeService struct {
}

func (ts TypeService) List() ([]model.Type, error) {
	Db := utils.GetDB()
	types := make([]model.Type, 4)
	if err := Db.Raw("SELECT id,type_name,amount FROM " + command.DBType + " ORDER BY amount DESC").Scan(&types).Error; err != nil {
		return nil, errors.New("查询失败")
	}

	return types, nil
}
