package contracts

type PaymentServices interface {
	HandlePayment(data []byte) error
	HandleLogging(data []byte) error
}
