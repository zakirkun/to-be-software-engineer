package contracts

import (
	"github.com/labstack/echo"
	"imzakir.dev/e-commerce/app/domains/models"
	"imzakir.dev/e-commerce/app/domains/types"
)

type CustomerRepository interface {
	Insert(data models.Customer) (*models.Customer, error)
	FindCustomerById(id int) (*models.Customer, error)
	
}

type CustomerServices interface {
	Insert(request types.RequestCreateCustomer) (*types.ResponseCreateCustomer, error)
	GetCustomerById(id int)  (*types.ResponseCreateCustomer, error)

}

type CustomerController interface {
	Insert(ctx echo.Context) error
	Show(ctx echo.Context) error
}
