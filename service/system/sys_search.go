package system

import "github.com/qianxia/blog/global"

type SearchService struct{}

/**
* 根据title搜索博客
 */
func (*SearchService) SearchBlog(title string) {
	global.RY_DB.Debug().Select("title").Where("title LIKE %?%", title)
}
