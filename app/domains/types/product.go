package types

import "imzakir.dev/e-commerce/app/domains/models"

type RequestCreateProduct struct {
	CategoryID         uint    `json:"category_id"`
	ProductName        string  `json:"product_name"`
	ProductImage       string  `json:"product_image"`
	ProductDescription string  `json:"product_description"`
	Price              float32 `json:"price"`
}

type ResponseCreateProduct struct {
	Product *models.Product `json:"product"`
}

type ResponsegetAllProduct struct {
	Product *[]models.Product `json:"product"`
}
