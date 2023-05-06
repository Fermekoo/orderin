package handler

import (
	"net/http"

	"github.com/Fermekoo/orderin-api/services"
	"github.com/Fermekoo/orderin-api/utils"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service *services.ProductService
}

func NewProductHandler(service *services.ProductService) *ProductHandler {
	return &ProductHandler{
		service: service,
	}
}

func (handler *ProductHandler) GetAll(ctx *gin.Context) {
	products, err := handler.service.Products()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(http.StatusInternalServerError, err))
		return
	}

	ctx.JSON(http.StatusOK, utils.ResponseOK(http.StatusOK, "products", products))
}
