package domains

import (
	"github.com/Fermekoo/orderin-api/db/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProductResponse struct {
	ID          uuid.UUID `json:"id"`
	CategoryID  uuid.UUID `json:"categoryId"`
	Category    string    `json:"category"`
	Title       string    `json:"title"`
	Price       uint64    `json:"price"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Size        uint64    `json:"size"`
	Color       string    `json:"color"`
}

type ProductService interface {
	Products(ctx *gin.Context) ([]ProductResponse, error)
	Product(productId uuid.UUID) (ProductResponse, error)
	ProductByCategory(categoryId string) ([]ProductResponse, error)
}

type ProductRepo interface {
	GetAll() ([]models.Product, error)
	FindById(productId uuid.UUID) (models.Product, error)
	GetProductByCategoryId(categoryId string) ([]models.Product, error)
}
