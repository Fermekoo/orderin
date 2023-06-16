package services

import (
	"github.com/Fermekoo/orderin-api/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductService struct {
	productRepo *repositories.ProductRepo
}

func NewProductService(db *gorm.DB) *ProductService {
	productRepo := repositories.NewProductRepo(db)

	return &ProductService{
		productRepo: productRepo,
	}
}

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

func (service *ProductService) Products() ([]ProductResponse, error) {
	var result = []ProductResponse{}
	products, err := service.productRepo.GetAll()
	if err != nil {
		return result, err
	}

	for _, p := range products {
		product := ProductResponse{
			ID:          p.ID,
			CategoryID:  p.CategoryID,
			Category:    p.Category.Category,
			Title:       p.Name,
			Price:       p.Price,
			Description: p.Description,
			Image:       p.Image,
			Size:        p.Size,
			Color:       p.Color,
		}

		result = append(result, product)
	}
	return result, nil
}

func (service *ProductService) Product(productId uuid.UUID) (ProductResponse, error) {
	var result = ProductResponse{}
	product, err := service.productRepo.FindById(productId)
	if err != nil {
		return result, err
	}

	result.ID = product.ID
	result.CategoryID = product.CategoryID
	result.Category = product.Category.Category
	result.Title = product.Name
	result.Price = product.Price
	result.Description = product.Description
	result.Image = product.Image

	return result, nil
}
