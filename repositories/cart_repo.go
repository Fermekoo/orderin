package repositories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CartRepo struct {
	db *gorm.DB
}

func NewCartRepo(db *gorm.DB) *CartRepo {
	return &CartRepo{
		db: db,
	}
}

type AddCart struct {
	UserID    uuid.UUID
	ProductID uuid.UUID
	Quantity  uint32
}

func (repo *CartRepo) Add(cart *Cart) error {
	err := repo.db.Create(&cart).Error

	return err
}
