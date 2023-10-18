package payment

type OrderPaymentStatus string

const (
	OrderPending OrderPaymentStatus = "pending"
	OrderSuccess OrderPaymentStatus = "success"
	OrderCancel  OrderPaymentStatus = "cancel"
	OrderExpired OrderPaymentStatus = "expired"
)

type PaymentVendor string

const (
	Midtrans PaymentVendor = "midtrans"
)
