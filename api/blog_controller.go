package api

import (
	db "Blog/db/sqlc"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type InsertBlogRequest struct {
	OwnerId int64  `json:"owner_id" binding:"required"`
	TypeId  int64  `json:"type_id" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	Image   string `json:"image" binding:"required"`
}

func (server *Server) insertBlog(ctx *gin.Context) {
	var req InsertBlogRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.SecureJSON(http.StatusBadRequest, err.Error)
		return
	}

	arg := db.InsertBlogParams{
		OwnerID: req.OwnerId,
		TypeID:  req.TypeId,
		Title:   req.Title,
		Content: req.Content,
		Image:   req.Image,
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

func (server *Server) incrViews(ctx *gin.Context) {

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.SecureJSON(http.StatusBadRequest, "invalid parameter")
		return
	}

	_, err = server.store.GetBlog(ctx, id)
	if err != nil {
		if err == ErrNoRows {
			ctx.SecureJSON(http.StatusNotFound, err.Error())
			return
		}
		ctx.SecureJSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = server.store.IncrViews(ctx, id)
	if err != nil {
		ctx.SecureJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.SecureJSON(http.StatusOK, "Increment Views Successfully")
}

func (server *Server) listBlogs(ctx *gin.Context) {

	blogs, err := server.store.ListBlogs(ctx)
	if err != nil {
		ctx.SecureJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.SecureJSON(http.StatusOK, blogs)
}

func (server *Server) getBlog(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	// if err != nil {
	// 	ctx.SecureJSON(http.StatusBadRequest, "invalid parameter")
	// 	return
	// }

	blog, err := server.store.GetBlog(ctx, id)
	if err != nil {
		if err == ErrNoRows {
			ctx.SecureJSON(http.StatusNotFound, err.Error())
			return
		}
		ctx.SecureJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.SecureJSON(http.StatusOK, blog)
}

func (server *Server) deleteBlog(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	// if err != nil {
	// 	ctx.SecureJSON(http.StatusBadRequest, "invalid parameter")
	// 	return
	// }

	err := server.store.DeleteBlog(ctx, id)
	if err != nil {
		ctx.SecureJSON(http.StatusInternalServerError, err)
		return
	}

	ctx.SecureJSON(http.StatusOK, "Delete Blog Successfully")
}

type UpdateBlogRequest struct {
	Id      int64   `json:"id" binding:"required"`
	TypId   *int64  `json:"type_id"`
	Title   *string `json:"title"`
	Content *string `json:"content"`
	Image   *string `json:"image"`
}

func (server *Server) updateBlog(ctx *gin.Context) {
	var req UpdateBlogRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.SecureJSON(http.StatusBadRequest, err.Error())
		return
	}

	arg := db.UpdateBlogParams{
		ID: req.Id,
	}

	if req.TypId != nil {
		arg.TypeID = sql.NullInt64{
			Int64: *req.TypId,
			Valid: true,
		}
	}

	if req.Title != nil {
		arg.Title = sql.NullString{
			String: *req.Title,
			Valid:  true,
		}
	}

	if req.Content != nil {
		arg.Content = sql.NullString{
			String: *req.Content,
			Valid:  true,
		}
	}

	if req.Image != nil {
		arg.Image = sql.NullString{
			String: *req.Image,
			Valid:  true,
		}
	}

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