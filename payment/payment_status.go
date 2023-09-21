package payment

type OrderPaymentStatus string

const (
	OrderPending OrderPaymentStatus = "pending"
	OrderSuccess OrderPaymentStatus = "success"
	OrderCancel  OrderPaymentStatus = "cancel"
	OrderExpired OrderPaymentStatus = "expired"
)
