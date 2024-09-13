package services

import (
	"golang.org/x/crypto/bcrypt"
	"imzakir.dev/e-commerce/app/domains/contracts"
	"imzakir.dev/e-commerce/app/domains/models"
	"imzakir.dev/e-commerce/app/domains/types"
	"imzakir.dev/e-commerce/app/repository"
)

type customerServices struct{}

var repo = repository.NewCustomerRepository()

func (c customerServices) AddCustomer(request types.RequestCreateCustomer) (*types.ResponseCreateCustomer, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	request.Password = string(hashedPassword)
	data, err := repo.Insert(models.Customer{
		Username: request.Username,
		Password: request.Password,
		Email:    request.Email,
		FullName: request.FullName,
	})
	if err != nil {
		return nil, err
	}

	return &types.ResponseCreateCustomer{
		Customer: data,
	}, nil
}

func NewCustomerServices() contracts.CustomerServices {
	return customerServices{}
}
