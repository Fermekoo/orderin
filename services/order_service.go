package services

import (
	"errors"
	"sync"

	"github.com/Fermekoo/orderin-api/db/models"
	"github.com/Fermekoo/orderin-api/domains"
	"github.com/Fermekoo/orderin-api/payment"
	"github.com/Fermekoo/orderin-api/repositories"
	"github.com/Fermekoo/orderin-api/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type orderService struct {
	config    utils.Config
	orderRepo domains.OrderRepo
	cartRepo  domains.CartRepo
}

func NewOrderService(config utils.Config, db *gorm.DB) domains.OrderService {
	orderRepo := repositories.NewOrderRepo(db)
	cartRepo := repositories.NewCartRepo(db)

	return &orderService{
		config:    config,
		orderRepo: orderRepo,
		cartRepo:  cartRepo,
	}
}

func (service *orderService) CreateInvoice(ctx *gin.Context, payloads domains.AddInvoice) error {
	authUser := getAuthUser(ctx)
	carts, err := service.cartRepo.GetSelectedItems(authUser.UserID, payloads.CartItems)
	if err != nil {
		return err
	}

	if len(carts) < 1 {
		return errors.New("your cart is empty")
	}

	var wgCarts sync.WaitGroup
	groupedMap := make(map[uuid.UUID]*models.Order)

	for _, c := range carts {
		wgCarts.Add(1)
		go func(c models.Cart) {
			defer wgCarts.Done()
			item := &models.OrderDetail{
				ID:        c.ID,
				ProductID: c.ProductID,
				Quantity:  c.Quantity,
				Price:     c.Product.Price,
				Total:     c.Product.Price * uint64(c.Quantity),
			}

			if merchant, exists := groupedMap[c.Product.Category.MerchantID]; exists {
				merchant.MerchantID = c.Product.Category.MerchantID
				merchant.Total += item.Total
				merchant.Fee = service.config.OrderFee
				merchant.TotalPayment = merchant.Total + merchant.Fee
				merchant.Details = append(merchant.Details, item)
				groupedMap[c.Product.Category.MerchantID] = merchant
			} else {
				groupedMap[c.Product.Category.MerchantID] = &models.Order{
					MerchantID:   c.Product.Category.MerchantID,
					Total:        item.Total,
					Fee:          service.config.OrderFee,
					TotalPayment: item.Total + service.config.OrderFee,
					Details:      []*models.OrderDetail{item},
				}
			}
		}(c)
	}

	wgCarts.Wait()

	var groupedInvoicesByMerchant []*models.Order

	var wgInvoice sync.WaitGroup

	errCh := make(chan error, 1)
	for _, merchant := range groupedMap {
		wgInvoice.Add(1)
		go func(merchant *models.Order) {
			defer wgInvoice.Done()
			orderID, err := uuid.NewRandom()
			if err != nil {
				errCh <- err
			}
			merchant.ID = orderID
			paymentVendor, err := payment.NewPayment(service.config)
			if err != nil {
				errCh <- err
			}

			payloadsPayment := &payment.CreatePayment{
				OrderID: merchant.ID,
				Bank:    payloads.PaymentChannel,
				Amount:  int(merchant.TotalPayment),
			}

			orderPayment, _ := paymentVendor.Pay(payloadsPayment)
			merchant.UserID = authUser.UserID
			paymentOrderId, err := uuid.NewRandom()
			if err != nil {
				errCh <- err
			}
			merchant.Payment = &models.PaymentOrder{
				ID:            paymentOrderId,
				OrderID:       orderID,
				Vendor:        orderPayment.PaymentVendor,
				Channel:       orderPayment.PaymentChannel,
				Total:         merchant.TotalPayment,
				PaymentFee:    0,
				PaymentStatus: payment.OrderPending,
				PaymentAction: orderPayment.PaymentAction,
				Type:          orderPayment.Type,
			}
			groupedInvoicesByMerchant = append(groupedInvoicesByMerchant, merchant)

		}(merchant)
	}
	wgInvoice.Wait()
	close(errCh)

	err = <-errCh
	if err != nil {
		return err
	}

	err = service.orderRepo.Create(groupedInvoicesByMerchant)

	return err
}
