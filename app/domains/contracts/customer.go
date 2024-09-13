package contracts

import (
	"github.com/labstack/echo"
	"imzakir.dev/e-commerce/app/domains/models"
	"imzakir.dev/e-commerce/app/domains/types"
)

type CustomerRepository interface {
	Insert(data models.Customer) (*models.Customer, error)
	//GetAll() (*[]models.Category, error)
	//Show(categoryId int) (*models.Category, error)
}

type CustomerServices interface {
	AddCustomer(request types.RequestCreateCustomer) (*types.ResponseCreateCustomer, error)
	//GetAll() (*types.ResponseListCategory, error)
	//Show(categoryId int) (*types.ResponseCreateCategory, error)
}

type CustomerController interface {
	Register(ctx echo.Context) error
	//GetAll(ctx echo.Context) error
	//Show(ctx echo.Context) error
}
