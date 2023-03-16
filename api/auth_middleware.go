package api

import (
	"Blog/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader     = "authorization"
	authorizationPayloadKey = "authorization_payload"
)

func AuthMiddlware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorization := ctx.Request.Header.Get(authorizationHeader)

		if len(authorization) == 0 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, " not authorization")
			return
		}

		payload, err := tokenMaker.VerifyToken(authorization)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
			return
		}

		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()
	}
}