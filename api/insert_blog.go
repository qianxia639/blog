package api

import (
	db "Blog/db/sqlc"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type insertBlogRequest struct {
	OwnerId int64  `json:"owner_id" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	Image   string `json:"image" binding:"required"`
}

func (server *Server) insertBlog(ctx *gin.Context) {
	var req insertBlogRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.SecureJSON(http.StatusBadRequest, err.Error())
		return
	}

	arg := db.InsertBlogParams{
		OwnerID:   req.OwnerId,
		Title:     req.Title,
		Content:   req.Content,
		Image:     req.Image,
		CreatedAt: time.Now(),
	}

	_, err := server.store.InsertBlog(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case ErrUniqueViolation:
				ctx.SecureJSON(http.StatusForbidden, err.Error())
				return
			}
		}
		ctx.SecureJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.SecureJSON(http.StatusOK, "Insert Blog Successful")
}
