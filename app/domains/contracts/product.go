package contracts

import (
	"github.com/labstack/echo"
	"imzakir.dev/e-commerce/app/domains/models"
	"imzakir.dev/e-commerce/app/domains/types"
)

type ProductRepository interface {
	Insert(data models.Product) (*models.Product, error)
	GetAll() ([]*models.Product, error)
	Get(productId int) (*models.Product, error)
	Update(productId int, data models.Product) (*models.Product, error)
	Delete(productId int) error
}

type ProductServices interface {
	Insert(request types.RequestProduct) (*types.ResponseCreateProduct, error)
	GetAll() (*types.ResponseListProduct, error)
	Get(productId int) (*types.ResponseProduct, error)
	Update(request types.RequestProduct, productId int) (*types.ResponseCreateProduct, error)
	Delete(productId int) error
}

type ProductController interface {
	Insert(ctx echo.Context) error
	GetAll(ctx echo.Context) error
	Get(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}
