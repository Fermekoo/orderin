package payment

import (
	"errors"
	"fmt"

	"github.com/Fermekoo/orderin-api/utils"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

type MidtransPayment struct {
}

var mdCore coreapi.Client

func NewMidtrans(config *utils.Config) Payment {
	var env midtrans.EnvironmentType
	if config.IS_PRODUCTION {
		env = midtrans.Production
	} else {
		env = midtrans.Sandbox
	}
	mdCore.New(config.MidtransServerKey, env)

	return &MidtransPayment{}
}

func (m *MidtransPayment) Pay(payloads *CreatePayment) (*ResponsePayment, error) {

	var result *ResponsePayment
	response, err := mdCore.ChargeTransaction(requestFormated(payloads))
	if err != nil {
		return result, err
	}

	return responseFormatted(response)
}

func (m *MidtransPayment) Inquiry(orderId string) (*ResponsePayment, error) {

	transaction, err := mdCore.CheckTransaction(orderId)
	if err != nil {
		return nil, err
	}

	var status OrderPaymentStatus
	if transaction.TransactionStatus == "capture" {
		if transaction.FraudStatus == "challenge" {
			status = OrderPending
		} else if transaction.FraudStatus == "accept" {
			status = OrderSuccess
		}
	} else if transaction.TransactionStatus == "settlement" {
		status = OrderSuccess
	} else if transaction.TransactionStatus == "deny" {
		status = OrderPending
	} else if transaction.TransactionStatus == "cancel" {
		status = OrderCancel
	} else if transaction.TransactionStatus == "expire" {
		status = OrderCancel
	} else {
		status = OrderPending
	}

	result := &ResponsePayment{
		TransactionID:   transaction.TransactionID,
		OrderID:         transaction.OrderID,
		PaymentVendor:   "midtrans",
		Status:          status,
		TransactionTime: transaction.TransactionTime,
	}

	return result, nil
}

func requestFormated(payloads *CreatePayment) *coreapi.ChargeReq {
	var chargeReq *coreapi.ChargeReq
	if payloads.Bank == "gopay" {
		expiry := 15 //minute
		chargeReq = &coreapi.ChargeReq{
			PaymentType: "gopay",
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  payloads.OrderID.String(),
				GrossAmt: int64(payloads.Amount),
			},
			Gopay: &coreapi.GopayDetails{
				EnableCallback: true,
				CallbackUrl:    "https://dandifermeko.com",
			},
			CustomExpiry: &coreapi.CustomExpiry{
				ExpiryDuration: expiry,
				Unit:           "minute",
			},
		}
	} else {
		switch payloads.Bank {
		case "mandiri":
			chargeReq = &coreapi.ChargeReq{
				PaymentType: coreapi.PaymentTypeEChannel,
				TransactionDetails: midtrans.TransactionDetails{
					OrderID:  payloads.OrderID.String(),
					GrossAmt: int64(payloads.Amount),
				},
				EChannel: &coreapi.EChannelDetail{
					BillInfo1: "payment with mandiri",
					BillInfo2: "mandiri midtrans",
					BillKey:   utils.RandomBillKey(),
				},
			}
		case "bca", "bri", "bni", "permata":
			chargeReq = &coreapi.ChargeReq{
				PaymentType: coreapi.PaymentTypeBankTransfer,
				TransactionDetails: midtrans.TransactionDetails{
					OrderID:  payloads.OrderID.String(),
					GrossAmt: int64(payloads.Amount),
				},
				BankTransfer: &coreapi.BankTransferDetails{
					Bank: midtrans.Bank(payloads.Bank),
				},
			}
		}
	}

	return chargeReq
}

func responseFormatted(response *coreapi.ChargeResponse) (*ResponsePayment, error) {
	var result ResponsePayment
	if response.StatusCode != "201" {
		return &result, errors.New(response.StatusMessage)
	}

	result.TransactionID = response.TransactionID
	result.OrderID = response.OrderID
	result.PaymentVendor = "midtrans"
	result.TransactionTime = response.TransactionTime
	result.Status = OrderPending

	switch response.PaymentType {
	case "bank_transfer":
		result.PaymentChannel = response.PaymentType
		if response.VaNumbers != nil {
			result.PaymentChannel = response.VaNumbers[0].Bank
			result.PaymentAction = response.VaNumbers[0].VANumber
		} else if response.PermataVaNumber != "" {
			result.PaymentAction = response.PermataVaNumber
			result.PaymentChannel = "permata"
		}
		result.Type = "va_number"
	case "echannel":
		result.PaymentChannel = "mandiri"
		result.PaymentAction = fmt.Sprintf("%s%s", response.BillerCode, response.BillKey)
	case "gopay", "shopeepay":
		result.PaymentChannel = "gopay"
		result.Type = "deeplink"
		result.PaymentAction = response.Actions[1].URL
	}

	return &result, nil

}
