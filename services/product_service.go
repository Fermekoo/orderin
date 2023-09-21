package services

import (
	"github.com/Fermekoo/orderin-api/db/models"
	"github.com/Fermekoo/orderin-api/domains"
	"github.com/Fermekoo/orderin-api/repositories"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type productService struct {
	productRepo domains.ProductRepo
}

func NewProductService(db *gorm.DB) domains.ProductService {
	productRepo := repositories.NewProductRepo(db)

	return &productService{
		productRepo: productRepo,
	}
}

func productResponses(products []models.Product) ([]domains.ProductResponse, error) {
	var result = []domains.ProductResponse{}

	for _, p := range products {
		product := domains.ProductResponse{
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
func (service *productService) Products(ctx *gin.Context) ([]domains.ProductResponse, error) {
	categoryId, CategoryIdExists := ctx.GetQuery("category")
	if CategoryIdExists && categoryId != "" {

		products, err := service.productRepo.GetProductByCategoryId(categoryId)
		if err != nil {
			return nil, err
		}
		return productResponses(products)
	} else {
		products, err := service.productRepo.GetAll()
		if err != nil {
			return nil, err
		}
		return productResponses(products)
	}
}

func (service *productService) Product(productId uuid.UUID) (domains.ProductResponse, error) {
	var result = domains.ProductResponse{}
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

func (service *productService) ProductByCategory(categoryId string) ([]domains.ProductResponse, error) {
	products, err := service.productRepo.GetProductByCategoryId(categoryId)

	if err != nil {
		return nil, err
	}

	return productResponses(products)
}
