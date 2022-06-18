package system

import (
	"log"

	"github.com/qianxia/blog/global"
	"github.com/qianxia/blog/model"
	"github.com/qianxia/blog/model/request"
	"gorm.io/gorm"
)

type CommentService struct{}

// 添加父级评论
func (*CommentService) ParentComment(pc request.ParentConment) (*model.Comment, error) {

	comment := &model.Comment{
		BlogId:   pc.BlogId,
		Nickname: pc.Nickname,
		Content:  pc.Content,
	}
	log.Printf("parent comment = %v\n", comment)
	return comment, global.QX_DB.Debug().Create(comment).Error
}

// 添加子级评论
func (*CommentService) ChildComment(cc request.ChildComment) (*model.Comment, error) {
	comment := &model.Comment{
		BlogId:   cc.BlogId,
		ParentId: cc.ParentId,
		Nickname: cc.Nickname,
		Content:  cc.Content,
	}

	log.Printf("child comment = %v\n", comment)
	return comment, global.QX_DB.Debug().Create(comment).Error
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
func (*CommentService) DeleteChildComment(parentId uint64) error {
	sql := `DELETE FROM qx_comment WHERE parent_id = ?`
	return global.QX_DB.Debug().Exec(sql, parentId).Error
}
