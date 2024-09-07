package types

import "imzakir.dev/e-commerce/app/domains/models"

type RequestCreateCategory struct {
	Name string `json:"category_name"`
}

type ResponseCreateCategory struct {
	Category *models.Category `json:"category"`
}

type ResponseListCategory struct {
	Category *[]models.Category `json:"category"`
}
