package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/Fermekoo/go-kapster/utils"
	"github.com/Fermekoo/go-kapster/utils/token"
	"github.com/gin-gonic/gin"
)

func JWTMiddleware(config utils.Config, tokenMaker token.TokenMaker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader(utils.AUTH_HEADER_KEY)
		if len(authHeader) < 1 {
			err := errors.New("access token is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse(http.StatusUnauthorized, err))
			return
		}

		fields := strings.Fields(authHeader)
		if len(fields) < 2 {
			err := errors.New("invalid access token format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse(http.StatusUnauthorized, err))
			return
		}

		if utils.AUTH_HEADER_TYPE != strings.ToLower(fields[0]) {
			err := fmt.Errorf("unsupported authorization type %s", utils.AUTH_HEADER_TYPE)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse(http.StatusUnauthorized, err))
			return
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(config.TokenSecretKey, accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse(http.StatusUnauthorized, err))
			return
		}

		ctx.Set(utils.AUTH_PAYLOAD_KEY, payload)
		ctx.Next()
	}
}
