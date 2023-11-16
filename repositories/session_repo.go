package repositories

import (
	"context"
	"fmt"

	"github.com/Fermekoo/orderin-api/db/models"
	"github.com/Fermekoo/orderin-api/domains"
	"gorm.io/gorm"
)

type sessionRepo struct {
	db *gorm.DB
}

func NewSessionRepo(db *gorm.DB) domains.SessionRepo {
	return &sessionRepo{
		db: db,
	}
}

func (repo *sessionRepo) Create(ctx context.Context, payload *models.Session) (models.Session, error) {
	err := repo.db.WithContext(ctx).Create(&payload).Error
	return *payload, err
}

func (repo *sessionRepo) FindByField(ctx context.Context, field string, value interface{}) (models.Session, error) {
	var session models.Session
	field = fmt.Sprintf("%s = ?", field)
	err := repo.db.WithContext(ctx).First(&session, field, value).Error

	return session, err
}
