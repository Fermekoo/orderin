package repositories

import (
	"context"

	"github.com/Fermekoo/orderin-api/db/models"
	"github.com/Fermekoo/orderin-api/domains"
	"github.com/Fermekoo/orderin-api/payment"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type orderRepo struct {
	db *gorm.DB
}

func NewOrderRepo(db *gorm.DB) domains.OrderRepo {
	return &orderRepo{
		db: db,
	}
}

func (repo *orderRepo) Create(ctx context.Context, checkout *models.Checkout) error {

	orders := checkout.Order
	err := repo.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		if err := tx.Omit("Order").Create(&checkout).Error; err != nil {
			return err
		}

		for _, invoice := range orders {
			orderId := uuid.New()
			order := &models.Order{
				ID:           orderId,
				CheckoutID:   checkout.ID,
				MerchantID:   invoice.MerchantID,
				Total:        invoice.Total,
				Fee:          invoice.Fee,
				TotalPayment: invoice.TotalPayment,
			}

			if err := tx.Create(order).Error; err != nil {
				return err //return any err will rollback
			}

			for _, detail := range invoice.Details {

				orderDetail := &models.OrderDetail{
					ID:        detail.ID,
					OrderID:   orderId,
					ProductID: detail.ProductID,
					Quantity:  detail.Quantity,
					Price:     detail.Price,
					Total:     detail.Total,
				}

				if err := tx.Create(orderDetail).Error; err != nil {
					return err
				}

				if err := tx.Table("products").Where("id = ?", detail.ProductID).Update("stock", gorm.Expr("stock - ?", detail.Quantity)).Error; err != nil {
					return err
				}

				if err := tx.Table("carts").Where("id = ?", detail.ID).Delete(nil).Error; err != nil {
					return err
				}
			}
		}

		return nil
	})

	return err
}

func (repo *orderRepo) GetCheckoutById(ctx context.Context, checkoutId uuid.UUID) (*models.Checkout, error) {
	var checkout *models.Checkout
	err := repo.db.WithContext(ctx).Where("id", checkoutId).First(&checkout).Error
	return checkout, err
}

func (repo *orderRepo) UpdateCheckoutStatus(ctx context.Context, updatePayload *domains.UpdateCheckout) error {

	err := repo.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Table("checkouts").Where("id = ?", updatePayload.CheckoutId).Where("payment_status", payment.OrderPending).Updates(map[string]interface{}{
			"payment_status": updatePayload.Status,
			"success_at":     updatePayload.SuccessAt,
		}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}
