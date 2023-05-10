package api

import (
	"Blog/core/errors"
	"Blog/core/logs"
	"Blog/core/result"
	db "Blog/db/sqlc"
	"Blog/utils"
	"net/http"
	"time"

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
		logs.Logs.Error(err)
		result.ParamError(ctx, errors.ParamErr.Error())
		return
	}

	hashPassword, err := utils.Encrypt(req.Password)
	if err != nil {
		logs.Logs.Error(err)
		result.ServerError(ctx, errors.ServerErr.Error())
		return
	}

	arg := &db.CreateUserParams{
		Username:     req.Username,
		Email:        req.Email,
		Nickname:     req.Email,
		Password:     hashPassword,
		RegisterTime: time.Now(),
	}

	_, err = server.store.CreateUser(ctx, arg)
	if err != nil {
		logs.Logs.Error(err)
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case errors.UniqueViolationErr:
				result.Error(ctx, http.StatusForbidden, errors.UsernameOrEmailEexistsErr.Error())
				return
			}
		}
		result.ServerError(ctx, errors.ServerErr.Error())
		return
	}

	result.OK(ctx, "Create Successfully")
}
