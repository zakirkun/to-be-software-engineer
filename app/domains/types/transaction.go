package types

import "teukufuad/e-commerce/app/domains/models"

type RequestOrder struct {
	ProductId int     `json:"product_id"`
	Qty       int     `json:"qty"`
	Amount    float64 `json:"amount"`
}

type ResponseOrder struct {
	Transaction *models.Transaction `json:"transaction"`
}
