package handler

import (
	"net/http"

	"github.com/Fermekoo/orderin-api/domains"
	"github.com/Fermekoo/orderin-api/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProductHandler struct {
	service domains.ProductService
}

func NewProductHandler(service domains.ProductService) *ProductHandler {
	return &ProductHandler{
		service: service,
	}
}

func (handler *ProductHandler) GetAll(ctx *gin.Context) {
	categoryId, _ := ctx.GetQuery("category")

	filter := domains.ProductSearch{
		Categories: &categoryId,
	}
	products, err := handler.service.Products(ctx, filter)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(http.StatusInternalServerError, err))
		return
	}

	ctx.JSON(http.StatusOK, utils.ResponseOK(http.StatusOK, "products", products))
}

func (handler *ProductHandler) Detail(ctx *gin.Context) {
	productId, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(http.StatusInternalServerError, err))
		return
	}

	product, err := handler.service.Product(ctx.Request.Context(), productId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(http.StatusInternalServerError, err))
		return
	}

	ctx.JSON(http.StatusOK, utils.ResponseOK(http.StatusOK, "product", product))
}
