package example

import (
	"errors"

	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/model"
	"github.com/qianxia/blog/model/response"
)

type ArchiveService struct{}

/**
* 按年份显示全部博客信息
 */
func (*ArchiveService) GetArchiveGroupByYear() (m map[string][]response.Archive, total int64, err error) {

	var archives []response.Archive

	var years []string
	m = make(map[string][]response.Archive)
	if err = global.QX_DB.Debug().Raw("SELECT FROM_UNIXTIME(updated_at, '%Y') AS year FROM qx_blog GROUP By year ORDER BY year DESC").Scan(&years).Error; err != nil {
		global.QX_LOG.Errorf("%s", err)
		return nil, 0, errors.New("失败1")
	}

	for _, year := range years {
		if err = global.QX_DB.Debug().Raw("SELECT id,title,updated_at,flag FROM qx_blog WHERE FROM_UNIXTIME(updated_at, '%Y') = ?", year).Scan(&archives).Error; err != nil {
			global.QX_LOG.Errorf("%s", err)
			return nil, 0, errors.New("失败2")
		}
		m[year] = archives
	}

	if err = global.QX_DB.Model(&model.Blog{}).Count(&total).Error; err != nil {
		global.QX_LOG.Errorf("%s", err)
		return nil, 0, errors.New("失败3")
	}

	return m, total, nil
}
