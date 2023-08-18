package repositories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderRepo struct {
	db *gorm.DB
}

func NewOrderRepo(db *gorm.DB) *OrderRepo {
	return &OrderRepo{
		db: db,
	}
}

type Invoice struct {
	OderID         uuid.UUID
	MerchantID     uuid.UUID
	UserID         uuid.UUID
	Total          uint64
	Fee            uint64
	TotalPayment   uint64
	PaymentAction  string
	PaymentChannel string
	PaymentResult  interface{}
	Details        []*InvoiceDetails
}

type InvoiceDetails struct {
	CartID    uuid.UUID
	ProductID uuid.UUID
	Quantity  uint32
	Price     uint64
	Total     uint64
}

func (repo *OrderRepo) Create(payloads []*Invoice) error {
	err := repo.db.Transaction(func(tx *gorm.DB) error {
		for _, invoice := range payloads {
			orderId, err := uuid.NewRandom()
			if err != nil {
				return err
			}
			order := &Order{
				ID:           orderId,
				UserID:       invoice.UserID,
				MerchantID:   invoice.MerchantID,
				Total:        invoice.Total,
				Fee:          invoice.Fee,
				TotalPayment: invoice.TotalPayment,
			}

			if err := tx.Create(&order).Error; err != nil {
				return err //return any err will rollback
			}

			for _, detail := range invoice.Details {
				orderDetailId, err := uuid.NewRandom()
				if err != nil {
					return err
				}
				orderDetail := &OrderDetail{
					ID:        orderDetailId,
					OrderID:   order.ID,
					ProductID: detail.ProductID,
					Quantity:  detail.Quantity,
					Price:     detail.Price,
					Total:     detail.Total,
				}

				if err := tx.Create(&orderDetail).Error; err != nil {
					return err
				}
				if err := tx.Table("products").Where("id = ?", detail.ProductID).Update("stock", gorm.Expr("stock - ?", detail.Quantity)).Error; err != nil {
					return err
				}

				if err := tx.Table("carts").Where("id = ?", detail.CartID).Delete(nil).Error; err != nil {
					return err
				}
			}
		}

		return nil
	})

	return err
}
