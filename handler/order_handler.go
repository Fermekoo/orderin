package handler

import (
	"net/http"

	"github.com/Fermekoo/orderin-api/services"
	"github.com/Fermekoo/orderin-api/utils"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	service *services.OrderService
}

func NewOrderHandler(service *services.OrderService) *OrderHandler {
	return &OrderHandler{
		service: service,
	}
}

func (handler *OrderHandler) Order(ctx *gin.Context) {
	var request services.AddInvoice
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(http.StatusBadRequest, err))
		return
	}

	err := handler.service.CreateInvoice(ctx, request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(http.StatusInternalServerError, err))
		return
	}
	ctx.JSON(http.StatusCreated, utils.ResponseOK(http.StatusCreated, "success", nil))
}
