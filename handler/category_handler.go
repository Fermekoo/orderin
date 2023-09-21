package handler

import (
	"net/http"

	"github.com/Fermekoo/orderin-api/domains"
	"github.com/Fermekoo/orderin-api/utils"
	"github.com/gin-gonic/gin"
)

type Categoryhandler struct {
	service domains.CategoryService
}

func NewCategoryHandler(service domains.CategoryService) *Categoryhandler {
	return &Categoryhandler{
		service: service,
	}
}

func (handler *Categoryhandler) GetAll(ctx *gin.Context) {
	categories, err := handler.service.Categories()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(http.StatusInternalServerError, err))
		return
	}

	ctx.JSON(http.StatusOK, utils.ResponseOK(http.StatusOK, "categories", categories))
}
