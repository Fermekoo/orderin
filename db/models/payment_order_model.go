package models

import (
	"time"

	"github.com/Fermekoo/orderin-api/payment"
	"github.com/google/uuid"
)

type PaymentOrder struct {
	ID            uuid.UUID `gorm:"primaryKey" json:"id"`
	OrderID       uuid.UUID
	Order         Order `gorm:"-,foreignKey:orderID"`
	Vendor        string
	Channel       string
	Total         uint64
	PaymentFee    uint64
	PaymentStatus payment.OrderPaymentStatus
	PaymentAction string
	Type          string
	SuccessAt     time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (PaymentOrder) TableName() string {
	return "payment_order"
}
