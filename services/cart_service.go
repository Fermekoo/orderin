package services

import (
	"errors"

	"github.com/Fermekoo/orderin-api/db/models"
	"github.com/Fermekoo/orderin-api/domains"
	"github.com/Fermekoo/orderin-api/repositories"
	"github.com/Fermekoo/orderin-api/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type cartService struct {
	config   *utils.Config
	cartRepo domains.CartRepo
}

func NewCartService(config *utils.Config, db *gorm.DB) domains.CartService {
	cartRepo := repositories.NewCartRepo(db)

	return &cartService{
		config:   config,
		cartRepo: cartRepo,
	}
}

func (service *cartService) Add(ctx *gin.Context, payload *domains.AddCart) error {
	authUser := getAuthUser(ctx)

	cart, err := service.cartRepo.FindByProductId(authUser.UserID, payload.ProductID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
		cartId, err := uuid.NewRandom()
		if err != nil {
			return err
		}

		cart := &models.Cart{
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
	}

	return err
}

func (service *cartService) GetAll(ctx *gin.Context) ([]domains.CartResponse, error) {
	var groupedMerchants []domains.CartResponse

	authUser := getAuthUser(ctx)
	carts, err := service.cartRepo.GetAll(authUser.UserID)
	if err != nil {
		return groupedMerchants, err
	}

	groupedMap := make(map[uuid.UUID]domains.CartResponse)
	for _, c := range carts {

		item := domains.CartProductResponse{
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
			groupedMap[c.Product.Category.MerchantID] = domains.CartResponse{
				MerchantId:   c.Product.Category.MerchantID,
				MerchantName: c.Product.Category.Merchant.Name,
				Items:        []domains.CartProductResponse{item},
				Total:        item.Total,
			}
		}
	}

	for _, merchant := range groupedMap {
		groupedMerchants = append(groupedMerchants, merchant)
	}

	return groupedMerchants, nil
}

func (service *cartService) UpdateQty(ctx *gin.Context, updateQty *domains.UpdateQty) error {
	authUser := getAuthUser(ctx)
	cartId, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		return err
	}

	err = service.cartRepo.UpdateQty(authUser.UserID, cartId, updateQty.Action)
	return err
}

func (service *cartService) Delete(ctx *gin.Context) error {
	authUser := getAuthUser(ctx)
	cartId, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		return err
	}
	err = service.cartRepo.Delete(authUser.UserID, cartId)

	return err
}
