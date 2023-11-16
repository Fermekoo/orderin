package domains

import (
	"context"

	"github.com/Fermekoo/orderin-api/db/models"
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
	Add(ctx context.Context, userId uuid.UUID, payload *AddCart) error
	GetAll(ctx context.Context, userId uuid.UUID) ([]CartResponse, error)
	UpdateQty(ctx context.Context, userId uuid.UUID, cartID uuid.UUID, updateQty *UpdateQty) error
	Delete(ctx context.Context, userId uuid.UUID, cartID uuid.UUID) error
}

type CartRepo interface {
	Add(ctx context.Context, cart *models.Cart) error
	GetAll(ctx context.Context, userId uuid.UUID) ([]models.Cart, error)
	UpdateQty(ctx context.Context, userId uuid.UUID, cartId uuid.UUID, act string) error
	Delete(ctx context.Context, userId uuid.UUID, cartId uuid.UUID) error
	FindByProductId(ctx context.Context, userId uuid.UUID, productId uuid.UUID) (models.Cart, error)
	GetSelectedItems(ctx context.Context, userId uuid.UUID, selectedIds []uuid.UUID) ([]models.Cart, error)
}
