package contracts

type PaymentServices interface {
	HandlePayment(data []byte) error
}
