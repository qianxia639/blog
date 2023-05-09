package api

import (
	"Blog/core/errors"
	"Blog/core/logs"
	"Blog/core/result"
	db "Blog/db/sqlc"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type insertArticleRequest struct {
	OwnerId    int64  `json:"owner_id" binding:"required"`
	Title      string `json:"title" binding:"required"`
	Content    string `json:"content" binding:"required"`
	Image      string `json:"image" binding:"required"`
	IsReward   bool   `json:"is_reward"`
	IsCritique bool   `json:"is_critique"`
}

func (server *Server) insertArticle(ctx *gin.Context) {
	var req insertArticleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		logs.Logs.Error(err)
		result.ParamError(ctx, errors.ParamErr.Error())
		return
	}

	arg := &db.InsertArticleParams{
		OwnerID:    req.OwnerId,
		Title:      req.Title,
		Content:    req.Content,
		Image:      req.Image,
		IsReward:   req.IsReward,
		IsCritique: req.IsCritique,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	_, err := server.store.InsertArticle(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case errors.UniqueViolationErr:
				logs.Logs.Error(err)
				result.Error(ctx, http.StatusForbidden, errors.TitleExistsErr.Error())
				return
			}
		}
		result.ServerError(ctx, errors.ServerErr.Error())
		return
	}

	result.OK(ctx, "Insert Article Successful")
}
