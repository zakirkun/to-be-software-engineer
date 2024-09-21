package contracts

import (
	"github.com/labstack/echo"
	"imzakir.dev/e-commerce/app/domains/models"
	"imzakir.dev/e-commerce/app/domains/types"
)

type CustomerRepository interface {
	Create(data models.Customer) error
	Update(id uint, data models.Customer) error
	GetWhere(where map[string]interface{}) (*models.Customer, error)
}
type CustomerServices interface {
	Register(request types.RequestRegisterCustomer) (*types.ResponseRegisterCustomer, error)
	Login(request types.RequestLoginCustomer) (*types.ResponseLoginCustomer, error)
}
type CustomerController interface {
	Register(ctx echo.Context) error
	Login(ctx echo.Context) error
}
