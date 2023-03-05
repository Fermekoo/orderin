package handler

import (
	"net/http"

	"github.com/Fermekoo/orderin-api/services"
	"github.com/Fermekoo/orderin-api/utils"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (handler *UserHandler) Register(ctx *gin.Context) {
	var request services.RegisterRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(http.StatusBadRequest, err))
		return
	}

	register, err := handler.service.Register(ctx, &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(http.StatusInternalServerError, err))
		return
	}

	ctx.JSON(http.StatusCreated, utils.ResponseOK(http.StatusCreated, "register success", register))
}

func (handler *UserHandler) Login(ctx *gin.Context) {
	var request services.LoginRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(http.StatusBadRequest, err))
		return
	}

	login, err := handler.service.Login(ctx, &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(http.StatusBadRequest, err))
		return
	}

	ctx.JSON(http.StatusOK, utils.ResponseOK(http.StatusOK, "login success", login))
}

func (handler *UserHandler) Profile(ctx *gin.Context) {
	profile, err := handler.service.Profile(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.ErrorResponse(http.StatusNotFound, err))
		return
	}

	ctx.JSON(http.StatusOK, utils.ResponseOK(http.StatusOK, "profile", profile))
}

func (handler *UserHandler) RefreshToken(ctx *gin.Context) {
	var request services.RenewAccessToken
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(http.StatusBadRequest, err))
		return
	}

	refreshToken, err := handler.service.RenewAccessToken(ctx, &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(http.StatusInternalServerError, err))
		return
	}

	ctx.JSON(http.StatusOK, utils.ResponseOK(http.StatusOK, "refresh token", refreshToken))
}
