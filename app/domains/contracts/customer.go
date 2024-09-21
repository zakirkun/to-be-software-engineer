package contracts

import (
	"github.com/labstack/echo"
	"imzakir.dev/e-commerce/app/domains/models"
	"imzakir.dev/e-commerce/app/domains/types"
)

type CustomerRepository interface {
	Insert(data models.Customer) (*models.Customer, error)
	GetByUsername(username string) (*models.Customer, error)
	GetWhere(where map[string]interface{}) (*models.Customer, error)
}

type CustomerServices interface {
	AddCustomer(request types.RequestCreateCustomer) (*types.ResponseCreateCustomer, error)
	Login(customer types.RequestLogin) (*types.ResponseLogin, error)
}

type CustomerController interface {
	Register(ctx echo.Context) error
	Login(ctx echo.Context) error
}
