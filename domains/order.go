package domains

import (
	"github.com/Fermekoo/orderin-api/db/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AddInvoice struct {
	CartItems      []uuid.UUID `json:"cartItems" binding:"required,dive"`
	PaymentChannel string      `json:"paymentChannel" binding:"required"`
}

type OrderService interface {
	CreateInvoice(ctx *gin.Context, payloads AddInvoice) error
}

type OrderRepo interface {
	Create(payloads *models.Checkout) error
}
