package handler

import (
	"net/http"

	"github.com/Fermekoo/orderin-api/domains"
	"github.com/Fermekoo/orderin-api/utils"
	"github.com/Fermekoo/orderin-api/utils/token"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service domains.UserService
}

func NewUserHandler(service domains.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (handler *UserHandler) Register(ctx *gin.Context) {
	var request domains.RegisterRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(http.StatusBadRequest, err))
		return
	}
	request.UserAgent = ctx.Request.UserAgent()
	request.IP = ctx.ClientIP()

	register, err := handler.service.Register(ctx.Request.Context(), &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(http.StatusInternalServerError, err))
		return
	}

	ctx.JSON(http.StatusCreated, utils.ResponseOK(http.StatusCreated, "register success", register))
}

func (handler *UserHandler) Login(ctx *gin.Context) {
	var request domains.LoginRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(http.StatusBadRequest, err))
		return
	}

	request.UserAgent = ctx.Request.UserAgent()
	request.IP = ctx.ClientIP()

	login, err := handler.service.Login(ctx.Request.Context(), &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(http.StatusBadRequest, err))
		return
	}

	ctx.JSON(http.StatusOK, utils.ResponseOK(http.StatusOK, "login success", login))
}

func (handler *UserHandler) Profile(ctx *gin.Context) {
	authUser := ctx.MustGet(utils.AUTH_PAYLOAD_KEY).(*token.Payload)
	profile, err := handler.service.Profile(ctx.Request.Context(), authUser.UserID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.ErrorResponse(http.StatusNotFound, err))
		return
	}

	ctx.JSON(http.StatusOK, utils.ResponseOK(http.StatusOK, "profile", profile))
}

func (handler *UserHandler) RefreshToken(ctx *gin.Context) {
	var request domains.RenewAccessToken
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(http.StatusBadRequest, err))
		return
	}

	refreshToken, err := handler.service.RenewAccessToken(ctx.Request.Context(), &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(http.StatusInternalServerError, err))
		return
	}

	ctx.JSON(http.StatusOK, utils.ResponseOK(http.StatusOK, "refresh token", refreshToken))
}
