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

/**
* 新增博客
 */
func (bs *BlogService) Save(saveBlog request.SaveBlog, userId uint64) error {

	// 根据userId查询用户信息
	var user model.User
	if err := global.DB.Debug().Where("id = ?", userId).First(&user).Error; err != nil {
		return err
	}
	// 根据post.tags[]的值查询对应的id
	tags := make([]model.Tag, 3)
	if err := global.DB.Debug().Where("tag_name in (?)", saveBlog.Tags).Find(&tags).Error; err != nil {
		return err
	}
	// 根据typeId查询
	var tp model.Type
	if err := global.DB.Debug().Model(&model.Type{}).Where("id = ?", saveBlog.TypeId).First(&tp).Error; err != nil {
		return err
	}

	// 构建数据
	// blog := model.Blog{
	// 	UserId:   userId,
	// 	Nickname: user.Nickname,
	// 	TypeId:   saveBlog.TypeId,
	// 	TypeName: tp.TypeName,
	// 	Title:    saveBlog.Title,
	// 	Content:  saveBlog.Content,
	// 	Flag:     saveBlog.Flag,
	// 	Tags:     tags,
	// }

	// 开启事务
	tx := global.DB.Begin()
	// 插入博客表数据以及博客标签中间表数据
	// if err := tx.Debug().Create(&blog).Error; err != nil {
	// 	tx.Rollback()
	// 	return err
	// }
	// // 更新分类表中amount字段的值
	// if err := tx.Model(&model.Type{Id: blog.TypeId}).Debug().Update("amount", gorm.Expr("amount + ?", 1)).Error; err != nil {
	// 	tx.Rollback()
	// 	return err
	// }

	// 提交事务
	return tx.Commit().Error
}

/**
* 个人博客列表展示
 */
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

/**
* 最新推荐展示
 */
func (bs *BlogService) LatestList() ([]model.Blog, error) {
	list := make([]model.Blog, 5)
	err := global.DB.Debug().Select("id,title").Order("updated_at DESC").Limit(5).Find(&list).Error

	return list, err
}

/**
* 首页博客展示及分页
 */
func (bs *BlogService) PageList(pageSize, pageNo int) (response.PageList, error) {
	var (
		total int64
		blogs []model.Blog
	)
	offset := (pageNo - 1) * pageSize
	err := global.DB.Debug().Preload("Tags").Offset(offset).Limit(pageSize).Find(&blogs).Count(&total).Error

	// 将分页信息和dataList封装到pageList中
	var pageList response.PageList
	pageList.Total = total
	pageList.PageNo = pageNo
	pageList.PageSize = pageSize

	pageList.DataList = blogs

	return pageList, err
}

/**
* 博客删除
 */
func (bs *BlogService) Delete(id uint64) error {
	var blog model.Blog
	if err := global.DB.Debug().Where("id = ?", id).First(&blog).Error; err != nil {
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

/**
* 修改博客
 */
func (*BlogService) Update(ub request.UpdateBlog) error {

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

/**
* 获取博客信息
 */
func (bs *BlogService) GetBlogInfo(id uint64) (*model.Blog, error) {

	var b model.Blog
	if err := global.DB.Debug().Preload("Tags").Where("id = ?", id).First(&b).Error; err != nil {
		return nil, err
	}

	// var user model.User
	// if err := global.DB.Debug().Where("id = ?", b.UserId).First(&user).Error; err != nil {
	// 	return nil, err
	// }

	// result := &response.BlogResult{
	// 	Id:        id,
	// 	Views:     b.Views,
	// 	Flag:      b.Flag,
	// 	Nickname:  b.Nickname,
	// 	Avatar:    user.Avatar,
	// 	TypeName:  b.TypeName,
	// 	Title:     b.Title,
	// 	Content:   b.Content,
	// 	UpdatedAt: b.UpdatedAt,
	// 	Tags:      b.Tags,
	// }

	// 返回
	return &b, nil
}

// 增加 博客的浏览次数
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
