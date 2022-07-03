package example

import (
	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/model"
	"github.com/qianxia/blog/model/response"
)

type ArchiveService struct{}

// @function GetArchiveGroupByYear
// @description 按年份显示全部博客信息
// @param {}
// @return map[string][]response.Archive, int64, error
func (*ArchiveService) GetArchiveGroupByYear() (m map[string][]response.Archive, total int64, err error) {

	var archives []response.Archive

	var years []string
	m = make(map[string][]response.Archive)
	sql := `SELECT FROM_UNIXTIME(updated_at, '%Y') AS year FROM t_blog GROUP By year ORDER BY year DESC`
	if err = global.DB.Debug().Raw(sql).Scan(&years).Error; err != nil {
		return
	}

	for _, year := range years {
		sql := `SELECT id,title,updated_at,flag FROM t_blog WHERE FROM_UNIXTIME(updated_at, '%Y') = ?`
		if err = global.DB.Debug().Raw(sql, year).Scan(&archives).Error; err != nil {
			return
		}
		m[year] = archives
	}

	if err = global.DB.Debug().Model(&model.Blog{}).Count(&total).Error; err != nil {
		return
	}

	return
}
