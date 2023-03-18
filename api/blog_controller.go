package api

import (
	db "Blog/db/sqlc"
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

type IncrVieswRequest struct {
	Id int64 `json:"id" binding:"required"`
}

func (server *Server) incrViews(ctx *gin.Context) {
	var req IncrVieswRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.SecureJSON(http.StatusBadRequest, err.Error())
		return
	}

	_, err := server.store.GetBlog(ctx, req.Id)
	if err != nil {
		if err == ErrNoRows {
			ctx.SecureJSON(http.StatusNotFound, err.Error())
			return
		}
		ctx.SecureJSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = server.store.IncrViews(ctx, req.Id)
	if err != nil {
		ctx.SecureJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.SecureJSON(http.StatusOK, "Increment Views Successful")
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

	err := server.store.DeleteBlog(ctx, id)
	if err != nil {
		ctx.SecureJSON(http.StatusInternalServerError, err)
		return
	}

	ctx.SecureJSON(http.StatusOK, "Delete Blog Successful")
}
