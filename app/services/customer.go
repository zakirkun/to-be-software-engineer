package services

import (
	"errors"

	"imzakir.dev/e-commerce/app/domains/contracts"
	"imzakir.dev/e-commerce/app/domains/models"
	"imzakir.dev/e-commerce/app/domains/types"
	"imzakir.dev/e-commerce/app/repository"
	"imzakir.dev/e-commerce/pkg/config"
	"imzakir.dev/e-commerce/pkg/jwt"
	"imzakir.dev/e-commerce/utils"
)

type customerServices struct{}

// Login implements contracts.CustomerServices.
func (c customerServices) Login(request types.RequestLoginCustomer) (*types.ResponseLoginCustomer, error) {

	where := make(map[string]interface{})
	where["username"] = request.Username

	repo := repository.NewCustomerRepository()
	getCust, err := repo.GetWhere(where)
	if err != nil {
		return nil, err
	}

	if getCust.ID == 0 {
		return nil, errors.New("record not found")
	}

	if ok := utils.CheckPasswordHash(request.Password, getCust.Password); !ok {
		return nil, errors.New("invalid password")
	}

	storeJwt := map[string]interface{}{
		"username": getCust.Username,
		"id":       getCust.ID,
	}

	// generate jwt
	_jwt := jwt.NewJWTImpl(config.GetString("jwt.signature_key"), config.GetInt("jwt.day_expired"))
	token, err := _jwt.GenerateToken(storeJwt)
	if err != nil {
		return nil, err
	}

	return &types.ResponseLoginCustomer{
		Token:    token,
		Username: getCust.Username,
	}, nil

}

// Register implements contracts.CustomerServices.
func (c customerServices) Register(request types.RequestRegisterCustomer) (*types.ResponseRegisterCustomer, error) {

	encPassword, _ := utils.HashPassword(request.Password)

	data := models.Customer{
		Username: request.Username,
		Password: encPassword,
		Email:    request.Email,
		FullName: request.FullName,
	}

	repo := repository.NewCustomerRepository()
	if err := repo.Create(data); err != nil {
		return nil, err
	}

	return &types.ResponseRegisterCustomer{
		Customer: data,
	}, nil

}

func NewCustomerServices() contracts.CustomerServices {
	return customerServices{}
}
