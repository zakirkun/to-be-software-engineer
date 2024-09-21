package services

import (

	"imzakir.dev/e-commerce/app/domains/contracts"
	"imzakir.dev/e-commerce/app/domains/models"
	"imzakir.dev/e-commerce/app/domains/types"
	"imzakir.dev/e-commerce/app/repository"
)


func NewCustomerServices() contracts.CustomerServices {
	return customerServices{}
}

type customerServices struct{}




// Insert implements contracts.CategoryServices.
func (c customerServices) Insert(request types.RequestCreateCustomer) (*types.ResponseCreateCustomer, error) {
	repo := repository.NewCustomerRepository()
	data, err := repo.Insert(models.Customer{
		Username: request.Username,
		Password: request.Password,
		FullName: request.FullName,
		Email: request.Email,
	})
	if err != nil {
		return nil, err
	}

	return &types.ResponseCreateCustomer{
		Customer: data,
	}, nil
}



// GetId implements contracts.CategoryServices.
func (c customerServices) GetCustomerById(id int) (*types.ResponseCreateCustomer, error) {
	data,err := repository.NewCustomerRepository().FindCustomerById(id)
	if err != nil {
		return nil, err
	}

	return &types.ResponseCreateCustomer{
		Customer: data,
	}, nil
}



