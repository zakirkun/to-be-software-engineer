package types

import "imzakir.dev/e-commerce/app/domains/models"

type RequestCreateCustomer struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
}

type ResponseCreateCustomer struct {
	Customer *models.Customer `json:"customer"`
}

//type ResponseListCategory struct {
//	Category *[]models.Category `json:"category"`
//}
