package models

import (
	"time"

	"github.com/google/uuid"
)

type Categories struct {
	ID         uuid.UUID `gorm:"primaryKey" json:"id"`
	MerchantID uuid.UUID
	Merchant   Merchant `gorm:"foreignKey:MerchantID"`
	Category   string
	Image      string
	IsEnable   bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (Categories) TableName() string {
	return "service_categories"
}
