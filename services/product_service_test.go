package services_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/Fermekoo/orderin-api/db/models"
	"github.com/Fermekoo/orderin-api/domains"
	"github.com/Fermekoo/orderin-api/domains/mocks"
	"github.com/Fermekoo/orderin-api/services"
	"github.com/Fermekoo/orderin-api/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetproductEmpty(t *testing.T) {
	productRepo := new(mocks.ProductRepo)
	productService := services.NewProductService(productRepo)
	productRepo.Mock.On("GetAll", context.Background()).Return([]models.Product{}, errors.New("empty products"))

	search := domains.ProductSearch{}
	products, err := productService.Products(context.Background(), search)
	assert.Error(t, err)
	assert.Empty(t, products)

	productRepo.AssertExpectations(t)

}

func TestProducts(t *testing.T) {
	productRepo := new(mocks.ProductRepo)
	productService := services.NewProductService(productRepo)

	var product models.Product
	mockProducts := []models.Product{}

	for i := 0; i < 10; i++ {
		productID, _ := uuid.NewRandom()
		categoryID, _ := uuid.NewRandom()
		product = models.Product{
			ID:         productID,
			CategoryID: categoryID,
			Category: models.Categories{
				ID: categoryID,
			},
			Name:        string(mock.AnythingOfType("string")),
			Price:       uint64(utils.RandomInt(1000, 10000)),
			Stock:       uint32(utils.RandomInt(2, 100)),
			IsEnable:    true,
			Description: string(mock.AnythingOfType("string")),
			Image:       string(mock.AnythingOfType("string")),
			Color:       string(mock.AnythingOfType("string")),
			Size:        uint64(utils.RandomInt(3, 10)),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}
		mockProducts = append(mockProducts, product)
	}
	productRepo.Mock.On("GetAll", context.Background()).Return(mockProducts, nil)

	products, err := productService.Products(context.Background(), domains.ProductSearch{})

	assert.NoError(t, err)
	assert.NotEmpty(t, products)
	assert.Len(t, products, 10)
	productRepo.AssertExpectations(t)
}

func TestProduct(t *testing.T) {
	productRepo := new(mocks.ProductRepo)
	productService := services.NewProductService(productRepo)

	productID, _ := uuid.NewRandom()
	categoryID, _ := uuid.NewRandom()
	product := models.Product{
		ID:         productID,
		CategoryID: categoryID,
		Category: models.Categories{
			ID: categoryID,
		},
		Name:        string(mock.AnythingOfType("string")),
		Price:       uint64(utils.RandomInt(1000, 10000)),
		Stock:       uint32(utils.RandomInt(2, 100)),
		IsEnable:    true,
		Description: string(mock.AnythingOfType("string")),
		Image:       string(mock.AnythingOfType("string")),
		Color:       string(mock.AnythingOfType("string")),
		Size:        uint64(utils.RandomInt(3, 10)),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	productRepo.Mock.On("FindById", context.Background(), productID).Return(product, nil).Once()

	productResponse, err := productService.Product(context.Background(), productID)

	assert.NoError(t, err)
	assert.NotEmpty(t, productResponse)
	assert.Equal(t, productID, productResponse.ID)
	productRepo.AssertExpectations(t)
}

func TestFilterProductByCategory(t *testing.T) {
	productRepo := new(mocks.ProductRepo)
	productService := services.NewProductService(productRepo)

	var product models.Product

	mockProducts := []models.Product{}

	categoryID, err := uuid.NewRandom()

	assert.NoError(t, err)

	for i := 0; i < 5; i++ {
		productID, err := uuid.NewRandom()
		assert.NoError(t, err)
		product = models.Product{
			ID:         productID,
			CategoryID: categoryID,
			Category: models.Categories{
				ID: categoryID,
			},
			Name:        string(mock.AnythingOfType("string")),
			Price:       uint64(utils.RandomInt(1000, 10000)),
			Stock:       uint32(utils.RandomInt(2, 100)),
			IsEnable:    true,
			Description: string(mock.AnythingOfType("string")),
			Image:       string(mock.AnythingOfType("string")),
			Color:       string(mock.AnythingOfType("string")),
			Size:        uint64(utils.RandomInt(3, 10)),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		mockProducts = append(mockProducts, product)
	}

	productRepo.Mock.On("GetProductByCategoryId", context.Background(), categoryID).Return(mockProducts, nil).Once()

	categoryIDstring := categoryID.String()
	search := domains.ProductSearch{
		Categories: &categoryIDstring,
	}

	products, err := productService.Products(context.Background(), search)
	assert.NoError(t, err)
	assert.NotEmpty(t, products)
	assert.Len(t, products, 5)

	for _, p := range products {
		assert.Equal(t, categoryID, p.CategoryID)
	}

	productRepo.AssertExpectations(t)

}

func TestProductNotFound(t *testing.T) {
	productRepo := new(mocks.ProductRepo)
	productService := services.NewProductService(productRepo)

	productID, _ := uuid.NewRandom()

	productRepo.Mock.On("FindById", context.Background(), productID).Return(models.Product{}, errors.New("data not found"))

	product, err := productService.Product(context.Background(), productID)

	assert.Error(t, err)
	assert.Empty(t, product)
	productRepo.AssertExpectations(t)
}

func TestProductByCategory(t *testing.T) {
	productRepo := new(mocks.ProductRepo)
	productService := services.NewProductService(productRepo)

	var product models.Product

	mockProducts := []models.Product{}

	categoryID, err := uuid.NewRandom()

	assert.NoError(t, err)

	for i := 0; i < 5; i++ {
		productID, err := uuid.NewRandom()
		assert.NoError(t, err)
		product = models.Product{
			ID:         productID,
			CategoryID: categoryID,
			Category: models.Categories{
				ID: categoryID,
			},
			Name:        string(mock.AnythingOfType("string")),
			Price:       uint64(utils.RandomInt(1000, 10000)),
			Stock:       uint32(utils.RandomInt(2, 100)),
			IsEnable:    true,
			Description: string(mock.AnythingOfType("string")),
			Image:       string(mock.AnythingOfType("string")),
			Color:       string(mock.AnythingOfType("string")),
			Size:        uint64(utils.RandomInt(3, 10)),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		mockProducts = append(mockProducts, product)
	}

	productRepo.Mock.On("GetProductByCategoryId", context.Background(), categoryID).Return(mockProducts, nil).Once()

	products, err := productService.ProductByCategory(context.Background(), categoryID)
	assert.NoError(t, err)
	assert.NotEmpty(t, products)
	assert.Len(t, products, 5)

	for _, p := range products {
		assert.Equal(t, categoryID, p.CategoryID)
	}

	productRepo.AssertExpectations(t)
}
