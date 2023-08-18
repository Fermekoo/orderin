package payment

import "github.com/google/uuid"

type CreatePayment struct {
	OrderID      uuid.UUID
	Bank         string
	Amount       int
	Name         string
	Email        string
	Phone        string
	Address      string
	RegisterDate string
}

type ResponsePayment struct {
	TransactionID   string `json:"transactionId"`
	OrderID         string `json:"orderId"`
	PaymentVendor   string `json:"paymentVendor"`
	PaymentType     string `json:"paymentType"`
	PaymentAction   string `json:"paymentAction"`
	Status          string `json:"status"`
	TransactionTime string `json:"transactionTime"`
}
