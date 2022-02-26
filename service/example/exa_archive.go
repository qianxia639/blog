package example

import (
	"errors"

	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/model"
	"github.com/qianxia/blog/response"
)

type ArchiveService struct{}

func (*ArchiveService) GetArchiveGroupByYear() (map[string][]response.Archive, int64, error) {

	var (
		total    int64
		archives []response.Archive
	)

	var years []string
	m := make(map[string][]response.Archive)
	if err := global.RY_DB.Debug().Raw("SELECT FROM_UNIXTIME(updated_at, '%Y') AS year FROM ry_blog GROUP By year ORDER BY year DESC").Scan(&years).Error; err != nil {
		global.RY_LOG.Errorf("%s", err)
		return nil, 0, errors.New("失败1")
	}

	for _, year := range years {
		if err := global.RY_DB.Debug().Raw("SELECT id,title,updated_at,flag FROM ry_blog WHERE FROM_UNIXTIME(updated_at, '%Y') = ?", year).Scan(&archives).Error; err != nil {
			global.RY_LOG.Errorf("%s", err)
			return nil, 0, errors.New("失败2")
		}
		m[year+"年"] = archives
	}

	if err := global.RY_DB.Model(&model.Blog{}).Count(&total).Error; err != nil {
		global.RY_LOG.Errorf("%s", err)
		return nil, 0, errors.New("失败3")
	}

	return m, total, nil
}
