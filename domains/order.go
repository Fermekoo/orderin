package domains

import (
	"context"
	"time"

	"github.com/Fermekoo/orderin-api/db/models"
	"github.com/Fermekoo/orderin-api/payment"
	"github.com/google/uuid"
)

type AddInvoice struct {
	CartItems      []uuid.UUID `json:"cartItems" binding:"required,dive"`
	PaymentChannel string      `json:"paymentChannel" binding:"required"`
}

type CallbackRequest struct {
	OrderID string `json:"order_id"`
}

type UpdateCheckout struct {
	CheckoutId uuid.UUID
	Status     payment.OrderPaymentStatus
	SuccessAt  time.Time
}
type OrderService interface {
	CreateInvoice(ctx context.Context, userID uuid.UUID, payloads *AddInvoice) error
	UpdateStatusPayment(ctx context.Context, orderId uuid.UUID) error
}

type OrderRepo interface {
	Create(ctx context.Context, payloads *models.Checkout) error
	GetCheckoutById(ctx context.Context, checkoutId uuid.UUID) (*models.Checkout, error)
	UpdateCheckoutStatus(ctx context.Context, updatePayload *UpdateCheckout) error
}
