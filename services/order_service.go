package services

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/Fermekoo/orderin-api/db/models"
	"github.com/Fermekoo/orderin-api/domains"
	"github.com/Fermekoo/orderin-api/payment"
	"github.com/Fermekoo/orderin-api/repositories"
	"github.com/Fermekoo/orderin-api/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type orderService struct {
	config    *utils.Config
	orderRepo domains.OrderRepo
	cartRepo  domains.CartRepo
}

func NewOrderService(config *utils.Config, db *gorm.DB) domains.OrderService {
	orderRepo := repositories.NewOrderRepo(db)
	cartRepo := repositories.NewCartRepo(db)

	return &orderService{
		config:    config,
		orderRepo: orderRepo,
		cartRepo:  cartRepo,
	}
}

func (service *orderService) CreateInvoice(ctx context.Context, userID uuid.UUID, payloads *domains.AddInvoice) error {

	carts, err := service.cartRepo.GetSelectedItems(ctx, userID, payloads.CartItems)
	if err != nil {
		return err
	}

	if len(carts) < 1 {
		return errors.New("your cart is empty")
	}

	var wgCarts sync.WaitGroup
	var mu sync.Mutex
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
				mu.Lock()
				merchant.Details = append(merchant.Details, item)
				mu.Unlock()
				groupedMap[c.Product.Category.MerchantID] = merchant
			} else {
				mu.Lock()
				groupedMap[c.Product.Category.MerchantID] = &models.Order{
					MerchantID:   c.Product.Category.MerchantID,
					Total:        item.Total,
					Fee:          service.config.OrderFee,
					TotalPayment: item.Total + service.config.OrderFee,
					Details:      []*models.OrderDetail{item},
				}
				mu.Unlock()
			}
		}(c)
	}

	wgCarts.Wait()

	var groupedInvoicesByMerchant []*models.Order
	var totalCheckout uint64
	var totalFeeCheckout uint64
	for _, merchant := range groupedMap {
		groupedInvoicesByMerchant = append(groupedInvoicesByMerchant, merchant)
		totalCheckout += merchant.Total
		totalFeeCheckout += totalFeeCheckout
	}

	totalPayment := totalCheckout + totalFeeCheckout
	paymentVendor, err := payment.NewPayment(service.config, payment.PaymentVendor(service.config.PaymentVendor))
	if err != nil {
		return err
	}

	checkoutId, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	paymentPayload := &payment.CreatePayment{
		OrderID: checkoutId,
		Bank:    payloads.PaymentChannel,
		Amount:  int(totalPayment),
	}

	createPayment, err := paymentVendor.Pay(paymentPayload)
	if err != nil {
		return err
	}

	checkout := &models.Checkout{
		ID:             checkoutId,
		UserID:         userID,
		Total:          totalPayment,
		PaymentVendor:  createPayment.PaymentVendor,
		PaymentChannel: createPayment.PaymentChannel,
		PaymentFee:     0,
		PaymentStatus:  payment.OrderPending,
		PaymentAction:  createPayment.PaymentAction,
		Type:           createPayment.Type,
		Order:          groupedInvoicesByMerchant,
	}

	err = service.orderRepo.Create(ctx, checkout)

	return err
}

func (service *orderService) UpdateStatusPayment(ctx context.Context, checkoutId uuid.UUID) error {

	checkout, err := service.orderRepo.GetCheckoutById(ctx, checkoutId)
	if err != nil {
		return err
	}

	if checkout.PaymentStatus != payment.OrderPending {
		return errors.New("invoice not pending")
	}

	pg, err := payment.NewPayment(service.config, payment.PaymentVendor(checkout.PaymentVendor))
	if err != nil {
		return err
	}

	transactions, err := pg.Inquiry(checkout.ID.String())
	if err != nil {
		return err
	}

	if transactions.Status == payment.OrderCancel || transactions.Status == payment.OrderSuccess {

		updatePayload := &domains.UpdateCheckout{
			CheckoutId: checkout.ID,
			Status:     transactions.Status,
		}
		if transactions.Status == payment.OrderSuccess {
			updatePayload.SuccessAt = time.Now()
		}

		err := service.orderRepo.UpdateCheckoutStatus(ctx, updatePayload)
		return err
	}

	return nil
}
