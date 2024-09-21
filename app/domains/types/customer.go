package types

import "imzakir.dev/e-commerce/app/domains/models"

type RequestCreateCustomer struct {
	Username string `json:"username"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
	Email string `json:"email"`


}

type ResponseCreateCustomer struct {
	Customer *models.Customer `json:"customer"`
}


type ResponseListCustomer struct {
	Customer *[]models.Customer `json:"customer"`
}
