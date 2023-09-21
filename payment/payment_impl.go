package payment

import (
	"errors"

	"github.com/Fermekoo/orderin-api/utils"
)

type PaymentImpl struct {
	Payment Payment
}

func NewPayment(config utils.Config) (*PaymentImpl, error) {
	var service Payment
	var err error
	switch config.PaymentVendor {
	case "midtrans":
		service = NewMidtrans(config)
	default:
		err = errors.New("service not available")
	}

	return &PaymentImpl{service}, err
}

func (p *PaymentImpl) Pay(payloads *CreatePayment) (*ResponsePayment, error) {
	return p.Payment.Pay(payloads)
}
