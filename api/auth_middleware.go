package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader     = "Authorization"
	authorizationPayloadKey = "Authorization_Payload"
)

func (server *Server) authMiddlware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorization := ctx.Request.Header.Get(authorizationHeader)

		if len(authorization) == 0 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthorized")
			return
		}

		payload, err := server.maker.VerifyToken(authorization)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
			return
		}

		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()
	}
}
