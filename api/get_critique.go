package api

import (
	"Blog/core/errors"
	"Blog/core/logs"
	"Blog/core/result"
	db "Blog/db/sqlc"
	"strconv"

	"github.com/gin-gonic/gin"
)

type critiqueTree struct {
	Critique  db.Critique    `json:"critique"`
	Childrens []*db.Critique `json:"childrens,omitempty"`
}

func (server *Server) getCritiques(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Query("id"), 10, 64)
	if err != nil {
		logs.Logs.Error(err)
		result.BadRequestError(ctx, errors.InvalidSyntaxErr.Error())
		return
	}

	critiques, err := server.store.GetCritiques(ctx, id)
	if err != nil {
		logs.Logs.Error(err)
		result.ServerError(ctx, errors.ServerErr.Error())
		return
	}

	// critiqueMap := make(map[int64][]*db.critique)
	// for _, critique := range critiques {
	// 	if critique.ParentID >= 0 {
	// 		critiqueMap[critique.ParentID] = append(critiqueMap[critique.ParentID], &critique)
	// 	}
	// }

	critiqueTreeMap := make(map[int64]*critiqueTree)
	for _, critique := range critiques {
		if critique.ParentID >= 0 {
			critiqueTreeMap[critique.ID] = &critiqueTree{
				Critique:  critique,
				Childrens: make([]*db.Critique, 0),
			}
		}
	}

	for _, critique := range critiques {
		// if critique.ParentID >= 0 {
		// 	parentID := critique.ParentID
		// 	if parentTree, ok := critiqueTreeMap[parentID]; ok {
		// 		parentTree.Replies = append(parentTree.Replies, &critique)
		// 	}
		// }
		if cs, err := server.store.GetChildCritiques(ctx, critique.ID); err == nil {
			for _, c := range cs {
				parentId := c.ParentID
				if parentTree, ok := critiqueTreeMap[parentId]; ok {
					parentTree.Childrens = append(parentTree.Childrens, &c)
				}
			}
		}
	}

	// critiqueTrees := make([]*critiqueTree, 0)
	// for _, tree := range critiqueTreeMap {
	// 	critiqueTrees = append(critiqueTrees, tree)
	// }

	result.Obj(ctx, critiqueTreeMap)
}
