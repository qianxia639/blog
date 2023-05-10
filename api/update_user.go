package api

import (
	"Blog/core/errors"
	"Blog/core/logs"
	"Blog/core/result"
	db "Blog/db/sqlc"
	"Blog/utils"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type updateUserRequest struct {
	Username string  `json:"username" binding:"required"`
	Email    *string `json:"email"`
	Nickname *string `json:"nickname"`
	Password *string `json:"password"`
	Avatar   *string `json:"avatar"`
}

func (server *Server) updateUser(ctx *gin.Context) {
	var req updateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		logs.Logs.Error(err)
		result.ParamError(ctx, errors.ParamErr.Error())
		return
	}

	payload, err := server.readToken(ctx.Request)
	if err != nil {
		result.UnauthorizedError(ctx, err.Error())
		return
	}

	if req.Username != payload.Username {
		result.UnauthorizedError(ctx, errors.UnauthorizedError.Error())
		return
	}

	arg := &db.UpdateUserParams{
		Username: req.Username,
		Email:    newNullString(req.Email),
		Nickname: newNullString(req.Nickname),
		Avatar:   newNullString(req.Avatar),
	}

	if req.Password != nil {
		hashPassword, err := utils.Encrypt(*req.Password)
		if err != nil {
			result.ServerError(ctx, errors.ServerErr.Error())
			return
		}
		arg.Password = sql.NullString{
			String: hashPassword,
			Valid:  true,
		}
	}

	user, err := server.store.UpdateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case errors.UniqueViolationErr:
				result.Error(ctx, http.StatusForbidden, errors.NicknameExistsErr.Error())
				return
			}
		}
		result.ServerError(ctx, errors.ServerErr.Error())
		return
	}

	key := fmt.Sprintf("t_%s", user.Username)
	err = server.rdb.Set(ctx, key, &user, 24*time.Hour).Err()
	if err != nil {
		logs.Logs.Error("redis err: ", err.Error())
		result.ServerError(ctx, errors.ServerErr.Error())
		return
	}

	result.OK(ctx, "Update Successfully")
}
