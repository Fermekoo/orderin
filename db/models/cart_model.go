package models

import (
	"time"

	"github.com/google/uuid"
)

type Cart struct {
	ID        uuid.UUID `gorm:"primaryKey" json:"id"`
	UserID    uuid.UUID
	ProductID uuid.UUID
	Product   Product `gorm:"foreignKey:ProductID"`
	Quantity  uint32
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Cart) TableName() string {
	return "carts"
}
