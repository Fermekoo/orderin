package repositories

import (
	"fmt"

	"gorm.io/gorm"
)

type SessionRepo struct {
	db *gorm.DB
}

func NewSessionRepo(db *gorm.DB) *SessionRepo {
	return &SessionRepo{
		db: db,
	}
}

func (repo *SessionRepo) Create(payload *Session) (Session, error) {
	err := repo.db.Create(&payload).Error
	return *payload, err
}

func (repo *SessionRepo) FindByField(field string, value interface{}) (Session, error) {
	var session Session
	field = fmt.Sprintf("%s = ?", field)
	err := repo.db.First(&session, field, value).Error

	return session, err
}
