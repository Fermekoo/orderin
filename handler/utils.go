package handler

import (
	"github.com/Fermekoo/orderin-api/utils"
	"github.com/Fermekoo/orderin-api/utils/token"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func getAuthUser(ctx *gin.Context) *token.Payload {
	authUser := ctx.MustGet(utils.AUTH_PAYLOAD_KEY).(*token.Payload)

	return authUser
}

func getUserId(ctx *gin.Context) uuid.UUID {
	auth := getAuthUser(ctx)

	return auth.UserID
}
