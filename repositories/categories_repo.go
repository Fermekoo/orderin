package repositories

import "gorm.io/gorm"

type CategoriesRepo struct {
	db *gorm.DB
}

func NewCategoriesRepo(db *gorm.DB) *CategoriesRepo {
	return &CategoriesRepo{
		db: db,
	}
}

func (repo *CategoriesRepo) GetAll() ([]Categories, error) {
	var categories []Categories

	err := repo.db.Preload("Merchant").Find(&categories).Error

	return categories, err
}
