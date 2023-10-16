package models

import (
	"time"

	"github.com/Fermekoo/orderin-api/payment"
	"github.com/google/uuid"
)

type Checkout struct {
	ID             uuid.UUID `gorm:"primaryKey" json:"id"`
	UserID         uuid.UUID
	Total          uint64
	PaymentVendor  string
	PaymentChannel string
	PaymentFee     uint64
	PaymentStatus  payment.OrderPaymentStatus
	PaymentAction  string
	Type           string
	Order          []*Order `json:",omitempty" gorm:"-,foreignKey:InvoiceID"`
	SuccessAt      time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
