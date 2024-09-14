package contracts

import (
	"github.com/labstack/echo"
	"imzakir.dev/e-commerce/app/domains/models"
	"imzakir.dev/e-commerce/app/domains/types"
)

type CategoryRepository interface {
	Insert(data models.Category) (*models.Category, error)
	GetAll() (*[]models.Category, error)
	GetDetail(id int) (*models.Category, error)
	Update(data models.Category) (*models.Category, error)
	Delete(data models.Category) (bool, error)
}

type CategoryServices interface {
	Insert(request types.RequestCreateCategory) (*types.ResponseCreateCategory, error)
	GetAll() (*types.ResponseListCategory, error)
	GetDetail(id int) (*types.ResponseGetDetailCategory, error)
	Update(id int, request types.RequestCreateCategory) (*types.ResponseCreateCategory, error)
	Delete(id int) (bool, error)
}

type CategoryController interface {
	Insert(ctx echo.Context) error
	GetAll(ctx echo.Context) error
	GetDetail(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}
