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
)

type updateBlogRequest struct {
	Id       int64   `json:"id" binding:"required"`
	Username string  `json:"username" binding:"required"`
	Title    *string `json:"title"`
	Content  *string `json:"content"`
	Image    *string `json:"image"`
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

func (server *Server) updateBlog(ctx *gin.Context) {
	var req updateBlogRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		logs.Logs.Error(err)
		result.BadRequestError(ctx, errors.ParamErr.Error())
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
		logs.Logs.Errorf("payload.Username: %s, user.Username: %s", payload.Username, user.Username)
		result.UnauthorizedError(ctx, errors.UnauthorizedError.Error())
		return
	}

	blog, err := server.store.GetBlog(ctx, req.Id)
	if err != nil {
		if err == ErrNoRows {
			logs.Logs.Error("Get Blog err: ", err)
			result.Error(ctx, http.StatusNotFound, errors.NotExistsUserErr.Error())
			return
		}
		logs.Logs.Error("Get Blog err: ", err)
		result.ServerError(ctx, errors.ServerErr.Error())
		return
	}

	if blog.OwnerID != user.ID {
		logs.Logs.Errorf("blog.OwnerID: %d, user.ID: %d", blog.OwnerID, user.ID)
		result.UnauthorizedError(ctx, errors.UnauthorizedError.Error())
		return
	}

	arg := &db.UpdateBlogParams{
		ID:        req.Id,
		UpdatedAt: time.Now(),
		Title:     newNullString(req.Title),
		Content:   newNullString(req.Content),
		Image:     newNullString(req.Image),
	}

	res, err := server.store.UpdateBlog(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case ErrUniqueViolation:
				result.Error(ctx, http.StatusForbidden, errors.TitleExistsErr.Error())
				return
			}
		}
		result.ServerError(ctx, errors.ServerErr.Error())
		return
	}

	result.Obj(ctx, res)
}
