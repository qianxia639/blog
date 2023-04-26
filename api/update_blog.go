package api

import (
	"Blog/core/errors"
	"Blog/core/result"
	db "Blog/db/sqlc"
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type updateBlogRequest struct {
	Id      int64   `json:"id" binding:"required"`
	Title   *string `json:"title"`
	Content *string `json:"content"`
	Image   *string `json:"image"`
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
		result.BadRequestError(ctx, errors.ParamErr.Error())
		return
	}

	arg := db.UpdateBlogParams{
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
				result.Error(ctx, http.StatusForbidden, err.Error())
				return
			}
		}
		result.ServerError(ctx, errors.ServerErr.Error())
		return
	}

	result.Obj(ctx, res)
}
