package models

import (
	"time"

	"github.com/google/uuid"
)

type OrderDetail struct {
	ID        uuid.UUID `gorm:"primaryKey" json:"id"`
	OrderID   uuid.UUID
	Order     Order `gorm:"foreignKey:OrderID" json:",omitempty"`
	ProductID uuid.UUID
	Product   Product `gorm:"foreignKey:ProductID" json:",omitempty"`
	Quantity  uint32
	Price     uint64
	Total     uint64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (OrderDetail) TableName() string {
	return "order_details"
}
