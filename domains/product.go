package domains

import (
	"context"

	"github.com/Fermekoo/orderin-api/db/models"
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

type ProductSearch struct {
	Categories *string `json:"category"`
}

type ProductService interface {
	Products(ctx context.Context, search ProductSearch) ([]ProductResponse, error)
	Product(ctx context.Context, productId uuid.UUID) (ProductResponse, error)
	ProductByCategory(cctx context.Context, ategoryId string) ([]ProductResponse, error)
}

type ProductRepo interface {
	GetAll(ctx context.Context) ([]models.Product, error)
	FindById(ctx context.Context, productId uuid.UUID) (models.Product, error)
	GetProductByCategoryId(ctx context.Context, categoryId string) ([]models.Product, error)
}
