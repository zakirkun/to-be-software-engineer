package types

import "imzakir.dev/e-commerce/app/domains/models"

type RequestCreateTransaction struct {
	ProductId int `json:"product_id"`
	CustomerId int `json:"customer_id"`
	Qty int `json:"qty"`
	Amount float64 `json:"amount"`
}

type ResponseCreateTransaction struct {
	Transaction *models.Transaction `json:"transaction"`
}

type ResponseListTransaction struct {
	Transaction *[]models.Transaction `json:"transaction"`
}
