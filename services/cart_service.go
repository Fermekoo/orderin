package services

import (
	"github.com/Fermekoo/orderin-api/repositories"
	"github.com/Fermekoo/orderin-api/utils"
	"github.com/Fermekoo/orderin-api/utils/token"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CartService struct {
	config   utils.Config
	cartRepo *repositories.CartRepo
}

func NewCartService(config utils.Config, db *gorm.DB) *CartService {
	cartRepo := repositories.NewCartRepo(db)

	return &CartService{
		config:   config,
		cartRepo: cartRepo,
	}
}

type AddCart struct {
	ProductID uuid.UUID `json:"productId" binding:"required"`
	Quantity  uint32    `json:"quantity" binding:"required"`
}

func (service *CartService) Add(ctx *gin.Context, payload *AddCart) error {
	authUser := ctx.MustGet(utils.AUTH_PAYLOAD_KEY).(*token.Payload)

	cartId, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	cart := &repositories.Cart{
		ID:        cartId,
		UserID:    authUser.UserID,
		ProductID: payload.ProductID,
		Quantity:  payload.Quantity,
	}

	err = service.cartRepo.Add(cart)
	return err
}
