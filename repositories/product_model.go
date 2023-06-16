package repositories

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID          uuid.UUID `gorm:"primaryKey" json:"id"`
	CategoryID  uuid.UUID
	Category    Categories `gorm:"foreignKey:CategoryID"`
	Name        string
	Price       uint64
	Stock       uint32
	IsEnable    bool
	Description string
	Image       string
	Color       string
	Size        uint64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (Product) TableName() string {
	return "products"
}
