package payment

import "github.com/google/uuid"

type Payment interface {
	Pay(payloads *CreatePayment) (interface{}, error)
	Inquiry(orderId uuid.UUID) (interface{}, error)
}
