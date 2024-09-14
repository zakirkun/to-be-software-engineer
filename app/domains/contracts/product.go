package contracts

import (
	"github.com/labstack/echo"
	"imzakir.dev/e-commerce/app/domains/models"
	"imzakir.dev/e-commerce/app/domains/types"
)

type ProductRepository interface {
	Insert(data models.Product) (*models.Product,error)
	GetAll() (*[]models.Product,error)
	FindProductById(id int) (*models.Product, error)
	DeleteById(id int) (*models.Product,error) 

}

type ProductServices interface {
	Insert(request types.RequestCreateProduct) (*types.ResponseCreateProduct, error)
	GetAll() (*types.ResponseListProduct,error)
	GetProductById(id int) (*types.ResponseCreateProduct, error)
	DeleteId(id int) (*types.ResponseCreateProduct, error)
}

type ProductController interface {
	Insert(ctx echo.Context) error
	GetAll(ctx echo.Context) error
	Show(ctx echo.Context) error
	Delete(ctx echo.Context) error
}