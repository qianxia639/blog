package api

import (
	"Blog/core/errors"
	"Blog/core/logs"
	"Blog/core/result"
	db "Blog/db/sqlc"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type commentTree struct {
	Comment db.Comment     `json:"comment"`
	Replies []*commentTree `json:"replies"`
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

	commentMap := make(map[int64][]*db.Comment)
	for _, comment := range comments {
		if comment.ParentID >= 0 {
			commentMap[comment.ParentID] = append(commentMap[comment.ParentID], &comment)
		}
	}

	commentTreeMap := make(map[int64]*commentTree)
	for _, comment := range comments {
		if comment.ParentID >= 0 {
			commentTreeMap[comment.ID] = &commentTree{
				Comment: comment,
				Replies: make([]*commentTree, 0),
			}
		}
		// server.store.GetChildComments(ctx, &db.GetChildCommentsParams{})
	}

	for _, comment := range comments {
		if comment.ParentID >= 0 {
			parentID := comment.ParentID
			if parentTree, ok := commentTreeMap[parentID]; ok {
				parentTree.Replies = append(parentTree.Replies, &commentTree{
					Comment: comment,
					Replies: make([]*commentTree, 0),
				})
			}
		}
	}

	commentTrees := make([]*commentTree, 0)
	for _, tree := range commentTreeMap {
		commentTrees = append(commentTrees, tree)
	}

	ctx.JSON(http.StatusOK, commentTrees)
}
