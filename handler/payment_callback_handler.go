package handler

import (
	"net/http"

	"github.com/Fermekoo/orderin-api/domains"
	"github.com/Fermekoo/orderin-api/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PaymentCallbackhandler struct {
	service domains.OrderService
}

func NewPaymentCallbackHandler(service domains.OrderService) *PaymentCallbackhandler {
	return &PaymentCallbackhandler{service: service}
}

func (handler *PaymentCallbackhandler) Callback(ctx *gin.Context) {
	var request domains.CallbackRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(http.StatusBadRequest, err))
		return
	}

	checkoutId, err := uuid.Parse(request.OrderID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(http.StatusInternalServerError, err))
		return
	}

	err = handler.service.UpdateStatusPayment(ctx, checkoutId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(http.StatusInternalServerError, err))
		return
	}

	ctx.JSON(http.StatusOK, utils.ResponseOK(http.StatusOK, "success", nil))
}
