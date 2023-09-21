package domains

import (
	"github.com/Fermekoo/orderin-api/db/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AddCart struct {
	ProductID uuid.UUID `json:"productId" binding:"required"`
	Quantity  uint32    `json:"quantity" binding:"required"`
}

type CartResponse struct {
	MerchantId   uuid.UUID             `json:"merchantId"`
	MerchantName string                `json:"merchantName"`
	Items        []CartProductResponse `json:"items"`
	Total        uint64                `json:"total"`
}

type CartProductResponse struct {
	CartId   uuid.UUID `json:"cartId"`
	Product  string    `json:"product"`
	Price    uint64    `json:"price"`
	Quantity uint32    `json:"quantity"`
	Total    uint64    `json:"total"`
}

type UpdateQty struct {
	Action string `json:"action" binding:"required,oneof=+ -"`
}

type CartService interface {
	Add(ctx *gin.Context, payload *AddCart) error
	GetAll(ctx *gin.Context) ([]CartResponse, error)
	UpdateQty(ctx *gin.Context, updateQty *UpdateQty) error
	Delete(ctx *gin.Context) error
}

type CartRepo interface {
	Add(cart *models.Cart) error
	GetAll(userId uuid.UUID) ([]models.Cart, error)
	UpdateQty(userId uuid.UUID, cartId uuid.UUID, act string) error
	Delete(userId uuid.UUID, cartId uuid.UUID) error
	FindByProductId(userId uuid.UUID, productId uuid.UUID) (models.Cart, error)
	GetSelectedItems(userId uuid.UUID, selectedIds []uuid.UUID) ([]models.Cart, error)
}
