package types

import "imzakir.dev/e-commerce/app/domains/models"

type RequestCreateCustomer struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
}

type RequestLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ResponseLogin struct {
	JwtToken string `json:"jwt_token"`
	//Message  string `json:"message"`
}

type ResponseCreateCustomer struct {
	Customer *models.Customer `json:"customer"`
}

//type ResponseListCategory struct {
//	Category *[]models.Category `json:"category"`
//}
