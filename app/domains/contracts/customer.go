package contracts

import (
	"github.com/labstack/echo"
	"imzakir.dev/e-commerce/app/domains/models"
	"imzakir.dev/e-commerce/app/domains/types"
)

type CustomerRepository interface {
	GetByEmail(email string) (*models.Customer, error)
	Insert(data models.Customer) (*models.Customer, error)
	Update(id int, data models.Customer) (*models.Customer, error)
}

type CustomerServices interface {
	SignIn(request types.RequestSignIn) (*types.ResponseSignIn, error)
	SignUp(request types.RequestSignUp) (*types.ResponseSignUp, error)
	//UpdateProfile() (*types.ResponseListCategory, error)
}

type CustomerController interface {
	SignIn(ctx echo.Context) error
	SignUp(ctx echo.Context) error
	//UpdateProfile(ctx echo.Context) error
}
