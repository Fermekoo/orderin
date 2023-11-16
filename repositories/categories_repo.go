package repositories

import (
	"context"

	"github.com/Fermekoo/orderin-api/db/models"
	"github.com/Fermekoo/orderin-api/domains"
	"gorm.io/gorm"
)

type categoriesRepo struct {
	db *gorm.DB
}

func NewCategoriesRepo(db *gorm.DB) domains.CategoriesRepo {
	return &categoriesRepo{
		db: db,
	}
}

func (repo *categoriesRepo) GetAll(ctx context.Context) ([]models.Categories, error) {
	var categories []models.Categories

	err := repo.db.WithContext(ctx).Find(&categories).Error

	return categories, err
}
