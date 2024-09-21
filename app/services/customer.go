package services

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"teukufuad/e-commerce/app/domains/contracts"
	"teukufuad/e-commerce/app/domains/models"
	"teukufuad/e-commerce/app/domains/types"
	"teukufuad/e-commerce/app/repository"
	"teukufuad/e-commerce/pkg/config"
	middelware "teukufuad/e-commerce/pkg/jwt"
)

type CustomerService struct {
}

func (c CustomerService) SignIn(request types.RequestSignIn) (*types.ResponseSignIn, error) {
	repo := repository.NewCustomerRepository()

	customer, err := repo.GetByEmail(request.Email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(request.Password)); err != nil {
		return nil, errors.New("username or password is incorrect")
	}

	jwt := middelware.NewJWTImpl(config.GetString("jwt.signature_key"), config.GetInt("jwt.day_expired"))
	token, err := jwt.GenerateToken(customer.ToMapCustomer())
	if err != nil {
		return nil, err
	}

	return &types.ResponseSignIn{
		Customer: customer,
		Token:    token,
	}, nil

}

func (c CustomerService) SignUp(request types.RequestSignUp) (*types.ResponseSignUp, error) {
	repo := repository.NewCustomerRepository()

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(request.Password), 8)

	data, err := repo.Insert(models.Customer{
		Username: request.Username,
		Password: string(hashedPassword),
		FullName: request.FullName,
		Email:    request.Email,
	})

	if err != nil {
		return nil, err
	}

	return &types.ResponseSignUp{
		Customer: data,
	}, nil
}

func NewCustomerService() contracts.CustomerServices {
	return CustomerService{}
}
