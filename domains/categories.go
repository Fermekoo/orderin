package domains

import (
	"github.com/Fermekoo/orderin-api/db/models"
	"github.com/google/uuid"
)

type CategoryResponse struct {
	ID         uuid.UUID `json:"id"`
	Category   string    `json:"category"`
	MerchantID uuid.UUID `json:"merchantId"`
	Merchant   string    `json:"merchant"`
	Image      string    `json:"image"`
}

type CategoryService interface {
	Categories() ([]CategoryResponse, error)
}

type CategoriesRepo interface {
	GetAll() ([]models.Categories, error)
}
