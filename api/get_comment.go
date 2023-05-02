package api

import (
	"Blog/core/errors"
	"Blog/core/logs"
	"Blog/core/result"
	db "Blog/db/sqlc"
	"strconv"

	"github.com/gin-gonic/gin"
)

type commentTree struct {
	Comment   db.Comment    `json:"comment"`
	Childrens []*db.Comment `json:"childrens,omitempty"`
}

func (server *Server) getComments(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Query("id"), 10, 64)
	if err != nil {
		logs.Logs.Error(err)
		result.BadRequestError(ctx, errors.InvalidSyntaxErr.Error())
		return
	}

	comments, err := server.store.GetComments(ctx, id)
	if err != nil {
		logs.Logs.Error(err)
		result.ServerError(ctx, errors.ServerErr.Error())
		return
	}

	// commentMap := make(map[int64][]*db.Comment)
	// for _, comment := range comments {
	// 	if comment.ParentID >= 0 {
	// 		commentMap[comment.ParentID] = append(commentMap[comment.ParentID], &comment)
	// 	}
	// }

	commentTreeMap := make(map[int64]*commentTree)
	for _, comment := range comments {
		if comment.ParentID >= 0 {
			commentTreeMap[comment.ID] = &commentTree{
				Comment:   comment,
				Childrens: make([]*db.Comment, 0),
			}
		}
	}

	for _, comment := range comments {
		// if comment.ParentID >= 0 {
		// 	parentID := comment.ParentID
		// 	if parentTree, ok := commentTreeMap[parentID]; ok {
		// 		parentTree.Replies = append(parentTree.Replies, &comment)
		// 	}
		// }
		if cs, err := server.store.GetChildComments(ctx, comment.ID); err == nil {
			for _, c := range cs {
				parentId := c.ParentID
				if parentTree, ok := commentTreeMap[parentId]; ok {
					parentTree.Childrens = append(parentTree.Childrens, &c)
				}
			}
		}
	}

	// commentTrees := make([]*commentTree, 0)
	// for _, tree := range commentTreeMap {
	// 	commentTrees = append(commentTrees, tree)
	// }

	result.Obj(ctx, commentTreeMap)
}
