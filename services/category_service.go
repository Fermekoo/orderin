package services

import (
	"github.com/Fermekoo/orderin-api/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CategoryService struct {
	categoriesRepo *repositories.CategoriesRepo
}

func NewCategoryService(db *gorm.DB) *CategoryService {
	categoriesRepo := repositories.NewCategoriesRepo(db)

	return &CategoryService{
		categoriesRepo: categoriesRepo,
	}
}

type CategoryResponse struct {
	ID         uuid.UUID `json:"id"`
	Category   string    `json:"category"`
	MerchantID uuid.UUID `json:"merchantId"`
	Merchant   string    `json:"merchant"`
	Image      string    `json:"image"`
}

func (service *CategoryService) Categories() ([]CategoryResponse, error) {
	var result = []CategoryResponse{}
	categories, err := service.categoriesRepo.GetAll()
	if err != nil {
		return result, err
	}

	for _, cat := range categories {
		category := CategoryResponse{
			ID:         cat.ID,
			Category:   cat.Category,
			MerchantID: cat.MerchantID,
			Merchant:   cat.Merchant.Name,
			Image:      cat.Image,
		}

		result = append(result, category)
	}

	return result, nil
}
