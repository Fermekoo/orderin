package repositories

import (
	"context"
	"fmt"

	"github.com/Fermekoo/orderin-api/db/models"
	"github.com/Fermekoo/orderin-api/domains"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) domains.UserRepo {
	return &userRepo{db: db}
}

func (repo *userRepo) Create(ctx context.Context, payload *models.User) (models.User, error) {
	err := repo.db.WithContext(ctx).Create(&payload).Error
	return *payload, err
}

func (repo *userRepo) FindByField(ctx context.Context, field string, value interface{}) (models.User, error) {
	var user models.User
	field = fmt.Sprintf("%s = ?", field)
	err := repo.db.WithContext(ctx).First(&user, field, value).Error

	return user, err
}
