package repositories

import (
	"context"

	"github.com/Fermekoo/orderin-api/db/models"
	"github.com/Fermekoo/orderin-api/domains"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type cartRepo struct {
	db *gorm.DB
}

func NewCartRepo(db *gorm.DB) domains.CartRepo {
	return &cartRepo{
		db: db,
	}
}

func (repo *cartRepo) Add(ctx context.Context, cart *models.Cart) error {
	err := repo.db.WithContext(ctx).Create(&cart).Error

	return err
}

func (repo *cartRepo) GetAll(ctx context.Context, userId uuid.UUID) ([]models.Cart, error) {
	var cart []models.Cart

	err := repo.db.WithContext(ctx).Preload("Product.Category.Merchant").Where("user_id =?", userId).Find(&cart).Error

	return cart, err
}

func (repo *cartRepo) UpdateQty(ctx context.Context, userId uuid.UUID, cartId uuid.UUID, act string) error {
	var cart models.Cart
	var query clause.Expr

	if act == "+" {
		query = gorm.Expr("quantity + ?", 1)
	} else {
		query = gorm.Expr("quantity - ?", 1)
	}

	err := repo.db.WithContext(ctx).Model(&cart).Where("user_id = ?", userId).Where("id = ?", cartId).Update("quantity", query).Error

	return err
}

func (repo *cartRepo) Delete(ctx context.Context, userId uuid.UUID, cartId uuid.UUID) error {
	var cart models.Cart

	err := repo.db.WithContext(ctx).Where("user_id = ?", userId).Where("id = ?", cartId).Delete(&cart).Error
	return err
}

func (repo *cartRepo) FindByProductId(ctx context.Context, userId uuid.UUID, productId uuid.UUID) (models.Cart, error) {
	var cart models.Cart
	err := repo.db.WithContext(ctx).Where("product_id", productId).Where("user_id", userId).First(&cart).Error

	return cart, err
}

func (repo *cartRepo) GetSelectedItems(ctx context.Context, userId uuid.UUID, selectedIds []uuid.UUID) ([]models.Cart, error) {
	var cart []models.Cart

	err := repo.db.WithContext(ctx).Preload("Product.Category.Merchant").Where("user_id", userId).Where("id in ?", selectedIds).Find(&cart).Error
	return cart, err
}
