package system

import (
	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/model"
	"github.com/qianxia/blog/model/request"
	"gorm.io/gorm"
)

type CommentService struct{}

// @function Save
// @description 添加评论
// @return *model.Comment, error
func (*CommentService) Save(comment request.Comment) (*model.Comment, error) {
	// 当parentId不等于0的时候，表示是子级评论，等于0的时候，表示该评论是父级评论
	c := &model.Comment{
		BlogId:   comment.BlogId,
		Nickname: comment.Nickname,
		Content:  comment.Content,
	}

	if comment.ParentId != 0 {
		c.ParentId = &comment.ParentId
	}

	return c, global.DB.Debug().Create(c).Error
}

// @function DeleteParentComment
// @description 删除父级评论
// @return error
func (*CommentService) DeleteParentComment(commentId uint64) error {
	sql := `DELETE FROM t_comment WHERE id = ?`
	sql1 := `DELETE FROM t_comment WHERE parent_id = ?`
	return global.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Debug().Exec(sql, commentId).Error; err != nil {
			return err
		}

		if err := tx.Debug().Exec(sql1, commentId).Error; err != nil {
			return err
		}
		return nil
	})
}

// @function DeleteChildComment
// @description 删除子级评论
// @return error
func (*CommentService) DeleteChildComment(id uint64) error {
	sql := `DELETE FROM t_comment WHERE id = ?`
	return global.DB.Debug().Exec(sql, id).Error
}

// @function List
// @description 评论列表
// @return []model.Comment, error
func (*CommentService) List(id uint64) ([]model.Comment, error) {
	var c []model.Comment
	err := global.DB.Debug().Where("blog_id = ?", id).Find(&c).Error
	return c, err
}
