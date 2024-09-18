package contracts

import (
	"github.com/labstack/echo"
	"github.com/morkid/paginate"
	"imzakir.dev/e-commerce/app/domains/models"
	"imzakir.dev/e-commerce/app/domains/types"
)

type ProductRepository interface {
	Create(data models.Product) error
	FindBy(where map[string]interface{}) (*[]models.Product, error)
	Update(id uint, data models.Product) error
	Delete(id uint) error
	GetAll() (*[]models.Product, error)
}
type ProductServices interface {
	Create(request types.RequestCreateProduct) (*types.ResponseCreateProduct, error)
	Delete(id uint) error
	GetAll() (*types.ResponsegetAllProduct, error)
	Update(id uint, request types.RequestCreateProduct) (*types.ResponseCreateProduct, error)
	GetByID(id uint) (*types.ResponsegetAllProduct, error)
	Pagination(ctx echo.Context) (*paginate.Page, error)
}
type ProductController interface {
	Insert(ctx echo.Context) error
	GetAll(ctx echo.Context) error
	GetByID(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
	Pagination(ctx echo.Context) error
}
