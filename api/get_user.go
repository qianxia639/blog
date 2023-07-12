package api

import (
	"Blog/core/errors"
	"Blog/core/logs"
	"Blog/core/result"
	db "Blog/db/sqlc"
	"Blog/utils"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type getUserResponse struct {
	Id           int64     `json:"id"`
	Usrename     string    `json:"username"`
	Email        string    `json:"email"`
	Nickname     string    `json:"nickname"`
	Avatar       string    `json:"avatar"`
	RegisterTime time.Time `json:"register_time"`
}

func (server *Server) getUser(ctx *gin.Context) {
	payload, err := server.readToken(ctx.Request)
	if err != nil {
		result.UnauthorizedError(ctx, err.Error())
		return
	}

	var user db.User
	if err := server.cache.Get(ctx, fmt.Sprintf("t_%s", payload.Username)).Scan(&user); err != nil {
		logs.Logger.Error("redis err: ", zap.Error(err))
		result.ServerError(ctx, errors.ServerErr.Error())
		return
	}

	resp := getUserResponse{
		Id:           user.ID,
		Usrename:     user.Username,
		Email:        utils.DesnsitizeEmail(user.Email),
		Nickname:     user.Nickname,
		Avatar:       user.Avatar,
		RegisterTime: user.RegisterTime,
	}

	result.Obj(ctx, resp)
}
