package types

import (
	"github.com/shopspring/decimal"
	"imzakir.dev/e-commerce/app/domains/models"
)

type RequestCreateTransaction struct {
	IdProduct int `json:"id_product"`
	//IdCustomer *int `json:"id_customer"`
	Qty    int             `json:"qty"`
	Amount decimal.Decimal `json:"amount"`
}

type ResponseCreateTransaction struct {
	Transaction *models.Transaction `json:"transaction"`
}
