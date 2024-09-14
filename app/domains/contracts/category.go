package contracts

import (
	"github.com/labstack/echo"
	"imzakir.dev/e-commerce/app/domains/models"
	"imzakir.dev/e-commerce/app/domains/types"
)

type CategoryRepository interface {
	Insert(data models.Category) (*models.Category, error)
	GetAll() (*[]models.Category, error)
	GetByID(id uint) (*models.Category, error)
	Update(id uint, data models.Category) (*models.Category, error)
	Delete(id uint) error
}

type CategoryServices interface {
	Insert(request types.RequestCreateCategory) (*types.ResponseCreateCategory, error)
	GetAll() (*types.ResponseListCategory, error)
	GetByID(id uint) (*types.ResponseCreateCategory, error)
	Update(id uint, request types.RequestUpdateCategory) (*types.ResponseUpdateCategory, error)
	Delete(id uint) error
}

type CategoryController interface {
	Insert(ctx echo.Context) error
	GetAll(ctx echo.Context) error
	GetByID(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}
