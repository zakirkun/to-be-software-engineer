package types

import "imzakir.dev/e-commerce/app/domains/models"

type RequestCreateProduct struct {
	NameProduct string `json:"product_name"`
	NameImage string `json:"product_image"`
	NameProduction string `json:"product_description"`
	NamePrice float64 `json:"price"`
	NameCategoryId int `json:"category_id"`
}

type ResponseCreateProduct struct {
	Product *models.Product `json:"product"`
}

type ResponseListProduct struct {
	Product *[]models.Product `json:"product"`
}