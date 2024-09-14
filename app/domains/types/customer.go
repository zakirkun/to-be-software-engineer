package types

import "imzakir.dev/e-commerce/app/domains/models"

type RequestSignUp struct {
	Username string `json:"username"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

type RequestSignIn struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

type ResponseSignUp struct {
	Customer *models.Customer `json:"customer"`
}

type ResponseSignIn struct {
	Customer *models.Customer `json:"customer"`
	Token    string
}
