package types

import "imzakir.dev/e-commerce/app/domains/models"

type RequestRegisterCustomer struct {
	Username string `json:"username"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

type ResponseRegisterCustomer struct {
	Customer models.Customer `json:"customer"`
}

type RequestLoginCustomer struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ResponseLoginCustomer struct {
	Token    string `json:"token"`
	Username string `json:"username"`
}
