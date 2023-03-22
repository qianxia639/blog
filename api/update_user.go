package api

import (
	db "Blog/db/sqlc"
	"Blog/token"
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
		ctx.SecureJSON(http.StatusBadRequest, err.Error())
		return
	}

	payload, ok := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if !ok {
		ctx.SecureJSON(http.StatusInternalServerError, "")
		return
	}

	if req.Username != payload.Username {
		ctx.SecureJSON(http.StatusBadRequest, "用户名错误")
		return
	}

	arg := db.UpdateUserParams{
		Username: req.Username,
	}

	if req.Email != nil {
		arg.Email = sql.NullString{
			String: *req.Email,
			Valid:  req.Email != nil,
		}
	}

	if req.Password != nil {
		hashPassword, err := utils.Encrypt(*req.Password)
		if err != nil {
			ctx.SecureJSON(http.StatusInternalServerError, err.Error())
			return
		}
		arg.Password = sql.NullString{
			String: hashPassword,
			Valid:  true,
		}
	}

	if req.Nickname != nil {
		arg.Nickname = sql.NullString{
			String: *req.Nickname,
			Valid:  req.Nickname != nil,
		}
	}

	if req.Avatar != nil {
		arg.Avatar = sql.NullString{
			String: *req.Avatar,
			Valid:  req.Avatar != nil,
		}
	}

	_, err := server.store.UpdateUser(ctx, arg)
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

	ctx.SecureJSON(http.StatusOK, "Update User Successfully")
}
