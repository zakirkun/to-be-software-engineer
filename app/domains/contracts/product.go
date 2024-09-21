package contracts

import (
	"github.com/labstack/echo"
	"github.com/morkid/paginate"
	"imzakir.dev/e-commerce/app/domains/models"
	"imzakir.dev/e-commerce/app/domains/types"
)

type ProductRepository interface {
	Insert(data models.Product) (*models.Product, error)
	FindAll() (*[]models.Product, error)
	FindById(productId int) (*models.Product, error)
	Update(data models.Product) (*models.Product, error)
	Delete(data models.Product) (bool, error)
}

type ProductServices interface {
	AddProduct(request types.RequestCreateProduct) (*types.ResponseCreateProduct, error)
	GetAllProducts() (*[]models.Product, error)
	GetDetail(productId int) (*models.Product, error)
	Update(productId int, request types.RequestCreateProduct) (*types.ResponseCreateProduct, error)
	Delete(productId int) (bool, error)
	Pagination(ctx echo.Context) (*paginate.Page, error)
}

type ProductController interface {
	Create(ctx echo.Context) error
	GetAll(ctx echo.Context) error
	GetDetail(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
	Pagination(ctx echo.Context) error
}
