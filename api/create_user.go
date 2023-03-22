package api

import (
	db "Blog/db/sqlc"
	"Blog/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type createUserRequest struct {
	Username string `json:"username" binding:"alphanum"`
	Email    string `json:"email" binding:"email"`
	Password string `json:"password" binding:"min=6,max=20"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.SecureJSON(http.StatusBadRequest, err.Error())
		return
	}

	hashPassword, err := utils.Encrypt(req.Password)
	if err != nil {
		ctx.SecureJSON(http.StatusInternalServerError, err.Error())
		return
	}

	arg := db.CreateUserParams{
		Username: req.Username,
		Email:    req.Email,
		Nickname: req.Email,
		Password: hashPassword,
	}

	_, err = server.store.CreateUser(ctx, arg)
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

	ctx.SecureJSON(http.StatusOK, "Create User Successfully")
}
