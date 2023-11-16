package services

import (
	"context"
	"errors"

	"github.com/Fermekoo/orderin-api/db/models"
	"github.com/Fermekoo/orderin-api/domains"
	"github.com/Fermekoo/orderin-api/repositories"
	"github.com/Fermekoo/orderin-api/utils"
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

func (service *cartService) Add(ctx context.Context, userId uuid.UUID, payload *domains.AddCart) error {

	cart, err := service.cartRepo.FindByProductId(ctx, userId, payload.ProductID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
		cartId, err := uuid.NewRandom()
		if err != nil {
			return err
		}

		cart := &models.Cart{
			ID:        cartId,
			UserID:    userId,
			ProductID: payload.ProductID,
			Quantity:  payload.Quantity,
		}
		err = service.cartRepo.Add(ctx, cart)
		if err != nil {
			return err
		}

	} else {
		err = service.cartRepo.UpdateQty(ctx, userId, cart.ID, "+")
	}

	return err
}

func (service *cartService) GetAll(ctx context.Context, userID uuid.UUID) ([]domains.CartResponse, error) {
	var groupedMerchants []domains.CartResponse

	carts, err := service.cartRepo.GetAll(ctx, userID)
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

func (service *cartService) UpdateQty(ctx context.Context, userID uuid.UUID, cartID uuid.UUID, updateQty *domains.UpdateQty) error {

	return service.cartRepo.UpdateQty(ctx, userID, cartID, updateQty.Action)
}

func (service *cartService) Delete(ctx context.Context, userID uuid.UUID, cartID uuid.UUID) error {

	return service.cartRepo.Delete(ctx, userID, cartID)
}
