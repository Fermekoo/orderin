package services

import (
	"context"

	"github.com/Fermekoo/orderin-api/domains"
)

type categoryService struct {
	categoriesRepo domains.CategoriesRepo
}

func NewCategoryService(categoriesRepo domains.CategoriesRepo) domains.CategoryService {
	return &categoryService{
		categoriesRepo: categoriesRepo,
	}
}

func (service *categoryService) Categories(ctx context.Context) ([]domains.CategoryResponse, error) {
	var result = []domains.CategoryResponse{}
	categories, err := service.categoriesRepo.GetAll(ctx)
	if err != nil {
		return result, err
	}

	for _, cat := range categories {
		category := domains.CategoryResponse{
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
