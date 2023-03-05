package repositories

import (
	"time"

	"github.com/google/uuid"
)

type Merchant struct {
	ID                uuid.UUID `gorm:"primaryKey" json:"id"`
	Name              string
	Phone             string
	Logo              string
	IsEnable          bool
	CreatedAt         time.Time
	UpdatedAt         time.Time
	ServiceCategories []Categories `gorm:"foreignKey:MerchantID"`
}

func (Merchant) TableName() string {
	return "merchants"
}
