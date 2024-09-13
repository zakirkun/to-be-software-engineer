package contracts

import (
	"github.com/labstack/echo"
	"imzakir.dev/e-commerce/app/domains/models"
	"imzakir.dev/e-commerce/app/domains/types"
)

type CategoryRepository interface {
	Insert(data models.Category) (*models.Category, error)
	GetAll() (*[]models.Category, error)
	Show(categoryId int) (*models.Category, error)
	Update(data *models.Category) (*models.Category, error)
}

type CategoryServices interface {
	Insert(request types.RequestCreateCategory) (*types.ResponseCreateCategory, error)
	GetAll() (*types.ResponseListCategory, error)
	Show(categoryId int) (*types.ResponseCreateCategory, error)
	Update(categoryId int, category types.RequestCreateCategory) (*types.ResponseCreateCategory, error)
}

type CategoryController interface {
	Insert(ctx echo.Context) error
	GetAll(ctx echo.Context) error
	Show(ctx echo.Context) error
	Edit(ctx echo.Context) error
}
