package services

import (
	"github.com/Fermekoo/orderin-api/utils"
	"github.com/Fermekoo/orderin-api/utils/token"
	"github.com/gin-gonic/gin"
)

func getAuthUser(ctx *gin.Context) *token.Payload {
	authUser := ctx.MustGet(utils.AUTH_PAYLOAD_KEY).(*token.Payload)

	return authUser
}
