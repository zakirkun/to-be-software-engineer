package contracts

import (
	"github.com/labstack/echo"
	"imzakir.dev/e-commerce/app/domains/models"
	"imzakir.dev/e-commerce/app/domains/types"
)

type ProductRepository interface {
	Insert(data models.Product) (*models.Product, error)
}

type ProductServices interface {
	AddProduct(request types.RequestCreateProduct) (*types.ResponseCreateProduct, error)
}

type ProductController interface {
	Create(ctx echo.Context) error
}
