package handler

import (
	"net/http"

	"github.com/Fermekoo/orderin-api/services"
	"github.com/Fermekoo/orderin-api/utils"
	"github.com/gin-gonic/gin"
)

type Categoryhandler struct {
	service *services.CategoryService
}

func NewCategoryHandler(service *services.CategoryService) *Categoryhandler {
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
