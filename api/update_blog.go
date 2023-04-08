package api

import (
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

func (server *Server) updateBlog(ctx *gin.Context) {
	var req updateBlogRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.SecureJSON(http.StatusBadRequest, err.Error())
		return
	}

	arg := db.UpdateBlogParams{
		ID:        req.Id,
		UpdatedAt: time.Now(),
	}

	for _, filed := range []struct {
		name  string
		value *string
	}{
		{"Title", req.Title},
		{"Content", req.Content},
		{"Image", req.Image},
	} {
		if filed.value != nil {
			argValue := sql.NullString{
				String: *filed.value,
				Valid:  true,
			}
			argValue.Scan(filed.value)
		}
	}

	// if req.Title != nil {
	// 	arg.Title = sql.NullString{
	// 		String: *req.Title,
	// 		Valid:  true,
	// 	}
	// }

	// if req.Content != nil {
	// 	arg.Content = sql.NullString{
	// 		String: *req.Content,
	// 		Valid:  true,
	// 	}
	// }

	// if req.Image != nil {
	// 	arg.Image = sql.NullString{
	// 		String: *req.Image,
	// 		Valid:  true,
	// 	}
	// }

	_, err := server.store.UpdateBlog(ctx, arg)
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

	ctx.SecureJSON(http.StatusOK, "Update Blog Successfully")
}
