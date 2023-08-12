package repositories

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID           uuid.UUID `gorm:"primaryKey" json:"id"`
	UserID       uuid.UUID
	MerchantID   uuid.UUID
	Total        uint64
	Fee          uint64
	TotalPayment uint64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (Order) TableName() string {
	return "orders"
}
