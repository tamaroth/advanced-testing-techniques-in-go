package mocks

type PaymentProvider interface {
	ProcessPayment(amount float64, currency string) (string, error)
}

type OrderService struct {
	paymentProvider PaymentProvider
}

func NewOrderService(provider PaymentProvider) *OrderService {
	return &OrderService{paymentProvider: provider}
}

// ProcessOrder attempts to process an order by charging the given amount.
func (o *OrderService) ProcessOrder(amount float64, currency string) (string, error) {
	return o.paymentProvider.ProcessPayment(amount, currency)
}
