package domains

import (
	"time"

	"github.com/Fermekoo/orderin-api/db/models"
	"github.com/Fermekoo/orderin-api/payment"
	"github.com/gin-gonic/gin"
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
	CreateInvoice(ctx *gin.Context, payloads AddInvoice) error
	UpdateStatusPayment(ctx *gin.Context, orderId uuid.UUID) error
}

type OrderRepo interface {
	Create(payloads *models.Checkout) error
	GetCheckoutById(checkoutId uuid.UUID) (*models.Checkout, error)
	UpdateCheckoutStatus(updatePayload *UpdateCheckout) error
}
