package types

import "imzakir.dev/e-commerce/app/domains/models"

type RequestCreateTransaction struct {
	ProductID int     `json:"product_id"`
	Username  string  `json:"-"`
	Qty       int     `json:"qty"`
	Amount    float32 `json:"amount"`
}

type ResponseCreateTransaction struct {
	TrxId int `json:"trx_id"`
}

type ResponseGetTransaction struct {
	Transaction *models.Transaction `json:"transaction"`
}
