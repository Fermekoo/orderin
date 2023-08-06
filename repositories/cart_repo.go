package repositories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func (repo *CartRepo) GetAll(userId uuid.UUID) ([]Cart, error) {
	var cart []Cart

	err := repo.db.Preload("Product").Where("user_id =?", userId).Find(&cart).Error

	return cart, err
}

func (repo *CartRepo) UpdateQty(userId uuid.UUID, cartId uuid.UUID, act string) error {
	var cart Cart
	var query clause.Expr

	if act == "+" {
		query = gorm.Expr("quantity + ?", 1)
	} else {
		query = gorm.Expr("quantity - ?", 1)
	}

	err := repo.db.Model(&cart).Where("user_id = ?", userId).Where("id = ?", cartId).Update("quantity", query).Error

	return err
}

func (repo *CartRepo) Delete(userId uuid.UUID, cartId uuid.UUID) error {
	var cart Cart

	err := repo.db.Where("user_id = ?", userId).Where("id = ?", cartId).Delete(&cart).Error
	return err
}
