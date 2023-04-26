package api

import (
	"Blog/core/errors"
	"Blog/core/result"
	"Blog/core/token"
	db "Blog/db/sqlc"
	"Blog/utils"
	"database/sql"
	"net/http"

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
		result.BadRequestError(ctx, errors.ParamErr.Error())
		return
	}

	// TODO: 这里的Key值后期需要更改
	payload, ok := ctx.MustGet("Authorization_Payload").(*token.Payload)
	if !ok {
		result.ServerError(ctx, "internal server error")
		return
	}

	if req.Username != payload.Username {
		result.BadRequestError(ctx, errors.UsernameErr.Error())
		return
	}

	arg := db.UpdateUserParams{
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

	_, err := server.store.UpdateUser(ctx, arg)
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

	result.OK(ctx, "Update User Successfully")
}
