package contracts

import (
	"github.com/labstack/echo"
	"imzakir.dev/e-commerce/app/domains/models"
	"imzakir.dev/e-commerce/app/domains/types"
)

type CategoryRepository interface {
	Insert(data models.Category) (*models.Category, error)
	GetAll() (*[]models.Category, error)
}

type CategoryServices interface {
	Insert(request types.RequestCreateCategory) (*types.ResponseCreateCategory, error)
	GetAll() (*types.ResponseListCategory, error)
}

type CategoryController interface {
	Insert(ctx echo.Context) error
	GetAll(ctx echo.Context) error
}
