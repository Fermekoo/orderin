package services

import (
	"log"

	"github.com/Fermekoo/orderin-api/repositories"
	"github.com/Fermekoo/orderin-api/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderService struct {
	config    utils.Config
	orderRepo *repositories.OrderRepo
	cartRepo  *repositories.CartRepo
}

func NewOrderService(config utils.Config, db *gorm.DB) *OrderService {
	orderRepo := repositories.NewOrderRepo(db)
	cartRepo := repositories.NewCartRepo(db)

	return &OrderService{
		config:    config,
		orderRepo: orderRepo,
		cartRepo:  cartRepo,
	}
}

type AddInvoice struct {
	CartItems []uuid.UUID `json:"cartItems" binding:"required,dive"`
}

func (service *OrderService) CreateInvoice(ctx *gin.Context, payloads AddInvoice) error {
	authUser := getAuthUser(ctx)
	carts, err := service.cartRepo.GetSelectedItems(authUser.UserID, payloads.CartItems)
	if err != nil {
		log.Fatal(err)
	}

	groupedMap := make(map[uuid.UUID]*repositories.Invoice)

	for _, c := range carts {
		item := repositories.InvoiceDetails{
			CartID:    c.ID,
			ProductID: c.ProductID,
			Quantity:  c.Quantity,
			Price:     c.Product.Price,
			Total:     c.Product.Price * uint64(c.Quantity),
		}

		if merchant, exists := groupedMap[c.Product.Category.MerchantID]; exists {
			merchant.MerchantID = c.Product.Category.MerchantID
			merchant.UserID = authUser.UserID
			merchant.Total += item.Total
			merchant.Fee = service.config.OderFee
			merchant.TotalPayment = merchant.Total + merchant.Fee
			merchant.Details = append(merchant.Details, &item)
			groupedMap[c.Product.Category.MerchantID] = merchant
		} else {
			groupedMap[c.Product.Category.MerchantID] = &repositories.Invoice{
				MerchantID:   c.Product.Category.MerchantID,
				UserID:       authUser.UserID,
				Total:        item.Total,
				Fee:          service.config.OderFee,
				TotalPayment: item.Total + service.config.OderFee,
				Details:      []*repositories.InvoiceDetails{&item},
			}
		}
	}

	var groupedInvoicesByMerchant []*repositories.Invoice
	for _, merchant := range groupedMap {
		groupedInvoicesByMerchant = append(groupedInvoicesByMerchant, merchant)
	}

	err = service.orderRepo.Create(groupedInvoicesByMerchant)
	return err
}
