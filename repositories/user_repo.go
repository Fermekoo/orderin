package repositories

import (
	"fmt"

	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (repo *UserRepo) Create(payload *User) (User, error) {
	err := repo.db.Create(&payload).Error
	return *payload, err
}

func (repo *UserRepo) FindByField(field string, value interface{}) (User, error) {
	var user User
	field = fmt.Sprintf("%s = ?", field)
	err := repo.db.First(&user, field, value).Error

	return user, err
}
