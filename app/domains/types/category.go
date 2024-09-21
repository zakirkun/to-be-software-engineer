package types

import "teukufuad/e-commerce/app/domains/models"

type RequestCreateCategory struct {
	Name string `json:"category_name"`
}

type ResponseCreateCategory struct {
	Category *models.Category `json:"category"`
}

type ResponseDeleteCategory struct {
	Message string `json:"message"`
}

type ResponseListCategory struct {
	Category *[]models.Category `json:"category"`
}

type ResponseCategory struct {
	Category *models.Category `json:"category"`
}
