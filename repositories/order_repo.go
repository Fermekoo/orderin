package repositories

import (
	"github.com/Fermekoo/orderin-api/db/models"
	"github.com/Fermekoo/orderin-api/domains"
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

func (repo *orderRepo) Create(payloads []*models.Order) error {
	err := repo.db.Transaction(func(tx *gorm.DB) error {
		for _, invoice := range payloads {

			order := &models.Order{
				ID:           invoice.ID,
				UserID:       invoice.UserID,
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
					OrderID:   invoice.ID,
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

			if err := tx.Create(&invoice.Payment).Error; err != nil {
				return err
			}

		}

		return nil
	})

	return err
}
