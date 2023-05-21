package api

import (
	"Blog/core/errors"
	"Blog/core/logs"
	"Blog/core/result"
	db "Blog/db/sqlc"
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"go.uber.org/zap"
)

type updateArticleRequest struct {
	Id         int64   `json:"id" binding:"required"`
	Username   string  `json:"username" binding:"required"`
	Title      *string `json:"title"`
	Content    *string `json:"content"`
	Image      *string `json:"image"`
	IsReward   *bool   `json:"is_reward"`
	IsCritique *bool   `json:"is_critique"`
}

func newNullString(s *string) sql.NullString {
	if s == nil {
		return sql.NullString{}
	}

	return sql.NullString{
		String: *s,
		Valid:  true,
	}
}

func (server *Server) updateArticle(ctx *gin.Context) {
	var req updateArticleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		logs.Logs.Error("bind pramar err", zap.Error(err))
		result.ParamError(ctx, errors.ParamErr.Error())
		return
	}

	authorization := ctx.GetHeader("Authorization")
	payload, err := server.maker.VerifyToken(authorization)
	if err != nil {
		result.UnauthorizedError(ctx, err.Error())
		return
	}

	user, _ := server.store.GetUser(ctx, req.Username)

	if payload.Username != user.Username {
		logs.Logs.Error("payload.Username not equals user.Username: %s", zap.String("payloadUsername", payload.Username), zap.String("userUsername", user.Username))
		result.UnauthorizedError(ctx, errors.UnauthorizedErr.Error())
		return
	}

	article, err := server.store.GetArticle(ctx, req.Id)
	if err != nil {
		if err == errors.NoRowsErr {
			logs.Logs.Error("Get Article err: ", zap.Error(err))
			result.Error(ctx, http.StatusNotFound, errors.NotExistsUserErr.Error())
			return
		}
		logs.Logs.Error("Get Article err: ", zap.Error(err))
		result.ServerError(ctx, errors.ServerErr.Error())
		return
	}

	if article.OwnerID != user.ID {
		logs.Logs.Error("article.OwnerID not equals user.ID: ", zap.Int64("articleOwnerID", article.OwnerID), zap.Int64("userID", user.ID))
		result.UnauthorizedError(ctx, errors.UnauthorizedErr.Error())
		return
	}

	arg := &db.UpdateArticleParams{
		ID:        req.Id,
		UpdatedAt: time.Now(),
		Title:     newNullString(req.Title),
		Content:   newNullString(req.Content),
		Image:     newNullString(req.Image),
	}

	res, err := server.store.UpdateArticle(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case errors.UniqueViolationErr:
				result.Error(ctx, http.StatusForbidden, errors.TitleExistsErr.Error())
				return
			}
		}
		result.ServerError(ctx, errors.ServerErr.Error())
		return
	}

	result.Obj(ctx, res)
}
