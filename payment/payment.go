package payment

import "github.com/google/uuid"

type Payment interface {
	Pay(payloads *CreatePayment) (*ResponsePayment, error)
	Inquiry(orderId uuid.UUID) (*ResponsePayment, error)
}
