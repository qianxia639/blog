package example

import (
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
	sql := `SELECT YEAR(updated_at, '%Y') AS year FROM qx_blog GROUP By year ORDER BY year DESC`
	if err = global.QX_DB.Debug().Raw(sql).Scan(&years).Error; err != nil {
		return nil, 0, nil
	}

	for _, year := range years {
		sql := `SELECT id,title,updated_at,flag FROM qx_blog WHERE YEAR(updated_at, '%Y') = ?`
		if err = global.QX_DB.Debug().Raw(sql, year).Scan(&archives).Error; err != nil {
			return nil, 0, nil
		}
		m[year] = archives
	}

	if err = global.QX_DB.Debug().Model(&model.Blog{}).Count(&total).Error; err != nil {
		return nil, 0, nil
	}

	return m, total, nil
}
