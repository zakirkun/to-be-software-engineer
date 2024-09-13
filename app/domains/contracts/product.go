package contracts

import (
	"github.com/labstack/echo"
	"imzakir.dev/e-commerce/app/domains/models"
	"imzakir.dev/e-commerce/app/domains/types"
)

type ProductRepository interface {
	Insert(data models.Product) (*models.Product, error)
	FindAll() (*[]models.Product, error)
	FindById(productId int) (*models.Product, error)
	Update(data models.Product) (*models.Product, error)
}

type ProductServices interface {
	AddProduct(request types.RequestCreateProduct) (*types.ResponseCreateProduct, error)
	GetAllProducts() (*[]models.Product, error)
	GetDetail(productId int) (*models.Product, error)
	Update(productId int, request types.RequestCreateProduct) (*types.ResponseCreateProduct, error)
}

type ProductController interface {
	Create(ctx echo.Context) error
	GetAll(ctx echo.Context) error
	GetDetail(ctx echo.Context) error
	Edit(ctx echo.Context) error
}
