package types

type PaymentParameter struct {
	Username  string `json:"username"`
	ProductId int    `json:"product_id"`
	TrxId     int    `json:"transaction_id"`
	Amount    int
}
