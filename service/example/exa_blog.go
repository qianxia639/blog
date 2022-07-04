package example

import (
	"errors"

	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/model"
	"github.com/qianxia/blog/model/request"
	"github.com/qianxia/blog/model/response"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BlogService struct{}

// @function Save
// @description 新增博客
// @return *model.Blog, error
func (bs *BlogService) SaveBlog(saveBlog request.SaveBlog, userId uint64) (*model.Blog, error) {

	// 根据userId查询用户信息
	var user model.User
	if err := global.DB.Debug().Where("id = ?", userId).First(&user).Error; err != nil {
		return nil, err
	}
	// 根据post.tags[]的值查询对应的id
	tags := make([]model.Tag, 3)
	if err := global.DB.Debug().Where("tag_name in (?)", saveBlog.Tags).Find(&tags).Error; err != nil {
		return nil, err
	}
	// 根据typeId查询
	var tp model.Type
	if err := global.DB.Debug().Model(&model.Type{}).Where("id = ?", saveBlog.TypeId).First(&tp).Error; err != nil {
		return nil, err
	}

	// 构建数据
	blog := model.Blog{
		UserId:   userId,
		User:     user,
		TypeId:   saveBlog.TypeId,
		TypeName: tp.TypeName,
		Title:    saveBlog.Title,
		Content:  saveBlog.Content,
		Flag:     saveBlog.Flag,
		Tags:     tags,
	}

	// 开启事务
	tx := global.DB.Begin()
	// 插入博客表数据以及博客标签中间表数据
	if err := tx.Debug().Create(&blog).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	// 更新分类表中amount字段的值
	if err := tx.Model(&model.Type{Id: blog.TypeId}).Debug().Update("amount", gorm.Expr("amount + ?", 1)).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// 提交事务
	return &blog, tx.Commit().Error
}

// @function List
// @description 个人博客列表展示
// @return *response.PageList, error
func (bs *BlogService) List(id uint64, pageNo, pageSize int) (*response.PageList, error) {
	var blogs []model.Blog
	var total int64

	offset := (pageNo - 1) * pageSize
	err := global.DB.Debug().Where("user_id = ?", id).Offset(offset).Limit(pageSize).Find(&blogs).Count(&total).Error

	var pageList response.PageList

	pageList.Total = total
	pageList.PageNo = pageNo
	pageList.PageSize = pageSize

	pageList.DataList = blogs

	return &pageList, err
}

// @function LatestList
// @description 最新推荐展示
// @return []model.Blog, error
func (bs *BlogService) LatestList() ([]model.Blog, error) {
	list := make([]model.Blog, 5)
	err := global.DB.Debug().Select("id,title").Order("updated_at DESC").Limit(5).Find(&list).Error

	return list, err
}

// @function PageList
// @description 首页博客展示及分页
// @return response.PageList, error
func (bs *BlogService) PageList(pageSize, pageNo int) (response.PageList, error) {
	var (
		total int64
		blogs []model.Blog
	)
	offset := (pageNo - 1) * pageSize
	err := global.DB.Debug().Preload("Tags").Preload("User").Offset(offset).Limit(pageSize).Find(&blogs).Count(&total).Error

	// 将分页信息和dataList封装到pageList中
	var pageList response.PageList
	pageList.Total = total
	pageList.PageNo = pageNo
	pageList.PageSize = pageSize

	pageList.DataList = blogs

	return pageList, err
}

// @function Delete
// @description 博客删除
// @return error
func (bs *BlogService) DeleteBlog(id, userId uint64) error {
	var blog model.Blog
	if err := global.DB.Debug().Where("id = ? AND user_id = ?", id, userId).First(&blog).Error; err != nil {
		return err
	}

	// 开启事务
	tx := global.DB.Begin()
	if err := tx.Debug().Select(clause.Associations).Delete(&blog).Error; err != nil {
		tx.Rollback() // // 事务回滚
		return err
	}

	if err := tx.Debug().Model(&model.Type{Id: blog.TypeId}).Update("amount", gorm.Expr("amount - ?", 1)).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 提交事务
	return tx.Commit().Error
}

// @function Update
// @description 修改博客
// @return error
func (*BlogService) UpdateBlog(ub request.UpdateBlog) error {

	blog := &model.Blog{
		Title:   ub.Title,
		Content: ub.Content,
		Flag:    ub.Flag,
	}

	err := global.DB.Debug().Where("id = ?", ub.Id).Updates(blog).Error
	if err != nil {
		return err
	}
	return nil
}

// @function GetBlogInfo
// @description 获取博客信息
// @return model.Blog, error
func (bs *BlogService) GetBlogInfo(id uint64) (blog model.Blog, err error) {
	err = global.DB.Debug().Preload("Tags").Preload("User").Where("id = ?", id).First(&blog).Error
	// 返回
	return
}

// @function IncrViews
// @description 增加博客的浏览次数
// @return error
func (bs *BlogService) IncrViews(id uint64) error {

	var blog = new(model.Blog)
	err := global.DB.Debug().Where("id = ?", id).First(blog).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if err := global.DB.Debug().Model(blog).UpdateColumn("views", gorm.Expr("views + 1")).Error; err != nil {
		return err
	}
	return nil
}

// @function QueryAll
// @description 博客列表(所有博客)
// @return response.PageList
func (bs *BlogService) QueryAll() (pageList response.PageList) {
	var pageNo = 1
	var pageSize = 20
	var offset = (pageNo - 1) * pageSize
	var blogs []model.Blog
	var totle int64
	global.DB.Debug().Model(&model.Blog{}).Limit(pageSize).Offset(offset).Find(&blogs).Count(&totle)

	pageList.PageNo = pageNo
	pageList.PageSize = pageSize
	pageList.Total = totle
	pageList.DataList = blogs

	return
}

// @function GetBlogGroupByFlag
// @description flag分组列表
// @return map[string][]model.Blog, error
func (bs *BlogService) GetBlogGroupByFlag() (map[string][]model.Blog, error) {

	var flags []string
	sql := `select flag from t_blog  GROUP BY flag`
	err := global.DB.Debug().Raw(sql).Scan(&flags).Error
	if err != nil {
		return nil, err
	}

	var blogs []model.Blog
	m := make(map[string][]model.Blog)
	for _, flag := range flags {
		sql := `select * from t_blog WHERE flag = ?`
		err = global.DB.Debug().Raw(sql, flag).Scan(&blogs).Error
		if err != nil {
			return nil, err
		}
		m[flag] = blogs
	}

	return m, nil
}
