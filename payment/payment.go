package payment

type Payment interface {
	Pay(payloads *CreatePayment) (*ResponsePayment, error)
	Inquiry(orderId string) (*ResponsePayment, error)
}
