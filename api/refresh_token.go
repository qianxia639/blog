package api

import (
	"Blog/core/errors"
	"Blog/core/logs"
	"Blog/core/result"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type refreshTokenRquest struct {
	Token string `json:"token"`
}

func (server *Server) refreshToken(ctx *gin.Context) {
	var req refreshTokenRquest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		logs.Logger.Error("Should Bing Body", zap.Error(err))
		result.ParamError(ctx, errors.ParamErr.Error())
		return
	}
	ua := ctx.Request.Header.Get("User-Agent")
	if ua == "" {
		logs.Logger.Error("Can't Find 'User-Agent' in header")
		result.ParamError(ctx, errors.ParamErr.Error())
		return
	}

	payload, err := server.maker.VerifyToken(req.Token)
	if err != nil {
		logs.Logger.Error("failed decode token", zap.Error(err))
		result.ParamError(ctx, errors.ParamErr.Error())
		return
	}

	token, _ := server.maker.CreateToken(payload.Username, server.conf.Token.AccessTokenDuration)

	result.Obj(ctx, token)

}
