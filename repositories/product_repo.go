package repositories

import "gorm.io/gorm"

type ProductRepo struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) *ProductRepo {
	return &ProductRepo{
		db: db,
	}
}

func (repo *ProductRepo) GetAll() ([]Product, error) {
	var products []Product

	err := repo.db.Preload("Category").Find(&products).Error

	return products, err
}
