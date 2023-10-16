package models

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID           uuid.UUID `gorm:"type:uuid,primaryKey;default:gen_random_uuid()" json:"id"`
	CheckoutID   uuid.UUID
	MerchantID   uuid.UUID
	Total        uint64
	Fee          uint64
	TotalPayment uint64
	Details      []*OrderDetail `json:",omitempty"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (Order) TableName() string {
	return "orders"
}
