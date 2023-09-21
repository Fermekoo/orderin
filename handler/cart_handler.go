package handler

import (
	"net/http"

	"github.com/Fermekoo/orderin-api/domains"
	"github.com/Fermekoo/orderin-api/utils"
	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	service domains.CartService
}

func NewCartHandler(service domains.CartService) *CartHandler {
	return &CartHandler{
		service: service,
	}
}

func (handler *CartHandler) AddCart(ctx *gin.Context) {
	var request domains.AddCart
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(http.StatusBadRequest, err))
		return
	}

	err := handler.service.Add(ctx, &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(http.StatusInternalServerError, err))
		return
	}

	ctx.JSON(http.StatusCreated, utils.ResponseOK(http.StatusCreated, "success", nil))
}

func (handler *CartHandler) GetAll(ctx *gin.Context) {
	carts, err := handler.service.GetAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(http.StatusInternalServerError, err))
		return
	}

	ctx.JSON(http.StatusOK, utils.ResponseOK(http.StatusOK, "carts user", carts))
}

func (handler *CartHandler) UpdateQty(ctx *gin.Context) {
	var request domains.UpdateQty
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(http.StatusBadRequest, err))
		return
	}

	err := handler.service.UpdateQty(ctx, &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(http.StatusInternalServerError, err))
		return
	}

	ctx.JSON(http.StatusOK, utils.ResponseOK(http.StatusOK, "success", nil))
}

func (handler *CartHandler) Delete(ctx *gin.Context) {
	err := handler.service.Delete(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, utils.ErrorResponse(http.StatusUnprocessableEntity, err))
		return
	}

	ctx.JSON(http.StatusOK, utils.ResponseOK(http.StatusOK, "success", nil))
}
