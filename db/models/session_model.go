package models

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID           uuid.UUID `gorm:"primaryKey" json:"id"`
	UserId       uuid.UUID
	RefreshToken string
	UserAgent    string
	ClientIP     string
	IsBlocked    bool
	ExpiresAt    time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (S Session) TableName() string {
	return "sessions"
}
