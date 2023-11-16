package repositories

import (
	"context"

	"github.com/Fermekoo/orderin-api/db/models"
	"github.com/Fermekoo/orderin-api/domains"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type productRepo struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) domains.ProductRepo {
	return &productRepo{
		db: db,
	}
}

func (repo *productRepo) GetAll(ctx context.Context) ([]models.Product, error) {
	var products []models.Product

	err := repo.db.WithContext(ctx).Preload("Category").Find(&products).Error

	return products, err
}

func (repo *productRepo) FindById(ctx context.Context, productId uuid.UUID) (models.Product, error) {
	var product models.Product

	err := repo.db.WithContext(ctx).Preload("Category").First(&product, "id = ?", productId).Error

	return product, err
}

func (repo *productRepo) GetProductByCategoryId(ctx context.Context, categoryId string) ([]models.Product, error) {
	var products []models.Product

	err := repo.db.WithContext(ctx).Preload("Category").Find(&products, "category_id = ?", categoryId).Error

	return products, err
}
