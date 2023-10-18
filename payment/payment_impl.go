package payment

import (
	"errors"

	"github.com/Fermekoo/orderin-api/utils"
)

type PaymentImpl struct {
	Payment       Payment
	paymentVendor PaymentVendor
}

func NewPayment(config *utils.Config, paymentVendor PaymentVendor) (*PaymentImpl, error) {
	var service Payment
	var err error
	switch paymentVendor {
	case "midtrans":
		service = NewMidtrans(config)
	default:
		err = errors.New("service not available")
	}

	return &PaymentImpl{
		Payment:       service,
		paymentVendor: paymentVendor,
	}, err
}

func (p *PaymentImpl) Pay(payloads *CreatePayment) (*ResponsePayment, error) {
	return p.Payment.Pay(payloads)
}

func (p *PaymentImpl) Inquiry(orderId string) (result *ResponsePayment, err error) {
	transaction, err := p.Payment.Inquiry(orderId)
	if err != nil {
		return nil, err
	}

	return transaction, err
}
