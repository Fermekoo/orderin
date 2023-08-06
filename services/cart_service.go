package services

import (
	"github.com/Fermekoo/orderin-api/repositories"
	"github.com/Fermekoo/orderin-api/utils"
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
	authUser := getAuthUser(ctx)

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

type CartResponse struct {
	CartId   uuid.UUID `json:"cartId"`
	Product  string    `json:"product"`
	Price    uint64    `json:"price"`
	Quantity uint32    `json:"quantity"`
}

func (service *CartService) GetAll(ctx *gin.Context) ([]CartResponse, error) {
	var result []CartResponse

	authUser := getAuthUser(ctx)
	carts, err := service.cartRepo.GetAll(authUser.UserID)
	if err != nil {
		return result, err
	}

	for _, c := range carts {
		cart := CartResponse{
			CartId:   c.ID,
			Product:  c.Product.Name,
			Price:    c.Product.Price,
			Quantity: c.Quantity,
		}

		result = append(result, cart)
	}

	return result, nil
}

type UpdateQty struct {
	Action string `json:"action" binding:"required,oneof=+ -"`
}

func (service *CartService) UpdateQty(ctx *gin.Context, updateQty *UpdateQty) error {
	authUser := getAuthUser(ctx)
	cartId, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		return err
	}

	err = service.cartRepo.UpdateQty(authUser.UserID, cartId, updateQty.Action)
	return err
}

func (service *CartService) Delete(ctx *gin.Context) error {
	authUser := getAuthUser(ctx)
	cartId, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		return err
	}
	err = service.cartRepo.Delete(authUser.UserID, cartId)

	return err
}
