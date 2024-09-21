package types

import "teukufuad/e-commerce/app/domains/models"

type RequestProduct struct {
	CategoryId  int     `json:"category_id"`
	Name        string  `json:"name"`
	Image       *string `json:"image"`
	Description *string `json:"Description"`
	Price       float64 `json:"price"`
}

type ResponseCreateProduct struct {
	Product *models.Product `json:"product"`
}

type ResponseDeleteProduct struct {
	Message string `json:"message"`
}

type ResponseListProduct struct {
	Product []*models.Product `json:"product"`
}

type ResponseProduct struct {
	Product *models.Product `json:"product"`
}
