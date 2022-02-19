package example

import (
	"errors"
	"fmt"

	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/model"
	"github.com/qianxia/blog/response"
	"github.com/qianxia/blog/utils"
)

type ArchiveService struct{}

func (as *ArchiveService) ArchivePageList(page map[string]int) (*response.PageList, error) {
	var (
		total    int64
		blogs    []model.Blog
		archives []response.Archive
	)
	if err := global.RY_DB.Debug().Model(&model.Blog{}).Select("id,title,updated_at,flag").Limit(page["pageSize"]).Offset(page["skipCount"]).Count(&total).Find(&blogs).Error; err != nil {
		return nil, errors.New("失败")
	}
	for _, v := range blogs {
		t := utils.TimestampToTime(v.UpdatedAt)
		ah := response.Archive{
			Id:    fmt.Sprintf("%v", v.Id),
			Title: v.Title,
			Flag:  v.Flag,
			Year:  fmt.Sprintf("%d年", t.Year()),
			Date:  fmt.Sprintf("%d月%d日", t.Month(), t.Day()),
		}
		archives = append(archives, ah)
	}
	var pageList response.PageList
	pageList.Total = total
	pageList.DataList = archives
	return &pageList, nil
}
