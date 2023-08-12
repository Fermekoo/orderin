package services

import (
	"errors"
	"fmt"

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

	cart, err := service.cartRepo.FindByProductId(authUser.UserID, payload.ProductID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
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
		if err != nil {
			return err
		}

	} else {
		err = service.cartRepo.UpdateQty(authUser.UserID, cart.ID, "+")
		fmt.Println("update")
	}

	return err
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

func (service *CartService) GetAll(ctx *gin.Context) ([]CartResponse, error) {
	var result []CartResponse

	authUser := getAuthUser(ctx)
	carts, err := service.cartRepo.GetAll(authUser.UserID)
	if err != nil {
		return result, err
	}

	groupedMap := make(map[uuid.UUID]CartResponse)
	for _, c := range carts {

		item := CartProductResponse{
			CartId:   c.ID,
			Product:  c.Product.Name,
			Price:    c.Product.Price,
			Quantity: c.Quantity,
			Total:    c.Product.Price * uint64(c.Quantity),
		}

		if merchant, exists := groupedMap[c.Product.Category.MerchantID]; exists {
			merchant.Total += item.Total
			merchant.Items = append(merchant.Items, item)        //add new item to merchant.Itemst
			groupedMap[c.Product.Category.MerchantID] = merchant // override old merchant with new Merchant which updated item
		} else {
			groupedMap[c.Product.Category.MerchantID] = CartResponse{
				MerchantId:   c.Product.Category.MerchantID,
				MerchantName: c.Product.Category.Merchant.Name,
				Items:        []CartProductResponse{item},
				Total:        item.Total,
			}
		}
	}

	var groupedMerchants []CartResponse
	for _, merchant := range groupedMap {
		groupedMerchants = append(groupedMerchants, merchant)
	}

	return groupedMerchants, nil
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
