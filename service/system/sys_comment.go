package system

import (
	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/model"
	"github.com/qianxia/blog/model/request"
	"gorm.io/gorm"
)

type CommentService struct{}

// 添加评论
func (*CommentService) Save(comment request.Comment) (*model.Comment, error) {
	// 当parentId不等于0的时候，表示是子级评论，等于0的时候，表示该评论是父级评论
	c := &model.Comment{
		BlogId:   comment.BlogId,
		Nickname: comment.Nickname,
		Content:  comment.Content,
	}

	if comment.ParentId != 0 {
		c.ParentId = comment.ParentId
	}

	return c, global.QX_DB.Debug().Create(c).Error
}

// 删除父级评论
func (*CommentService) DeleteParentComment(commentId uint64) error {
	sql := `DELETE FROM qx_comment WHERE id = ?`
	sql1 := `DELETE FROM qx_comment WHERE parent_id = ?`
	return global.QX_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Debug().Exec(sql, commentId).Error; err != nil {
			return err
		}

		if err := tx.Debug().Exec(sql1, commentId).Error; err != nil {
			return err
		}
		return nil
	})
}

// 删除子级评论
func (*CommentService) DeleteChildComment(id uint64) error {
	sql := `DELETE FROM qx_comment WHERE id = ?`
	return global.QX_DB.Debug().Exec(sql, id).Error
}
