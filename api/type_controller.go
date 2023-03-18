package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type InsertTypeRequest struct {
	TypeName string `json:"type_name" binding:"required"`
}

func (server *Server) insertType(ctx *gin.Context) {
	var req InsertTypeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.SecureJSON(http.StatusBadRequest, err.Error)
		return
	}

	_, err := server.store.InsertType(ctx, req.TypeName)
	if err != nil {
		if pqEerr, ok := err.(*pq.Error); ok {
			switch pqEerr.Code.Name() {
			case ErrUniqueViolation:
				ctx.SecureJSON(http.StatusForbidden, err.Error())
				return
			}
		}
		ctx.SecureJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.SecureJSON(http.StatusOK, "Insert Type Successfully")
}
