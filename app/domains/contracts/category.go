package contracts

import (
	"github.com/labstack/echo"
	"teukufuad/e-commerce/app/domains/models"
	"teukufuad/e-commerce/app/domains/types"
)

type CategoryRepository interface {
	Insert(data models.Category) (*models.Category, error)
	GetAll() (*[]models.Category, error)
	Get(categoryId int) (*models.Category, error)
	Update(categoryId int, data models.Category) (*models.Category, error)
	Delete(categoryId int) error
}

type CategoryServices interface {
	Insert(request types.RequestCreateCategory) (*types.ResponseCreateCategory, error)
	GetAll() (*types.ResponseListCategory, error)
	Get(categoryId int) (*types.ResponseCategory, error)
	Update(request types.RequestCreateCategory, categoryId int) (*types.ResponseCreateCategory, error)
	Delete(categoryId int) error
}

type CategoryController interface {
	Insert(ctx echo.Context) error
	GetAll(ctx echo.Context) error
	Get(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
	GetByCategory(ctx echo.Context) error
	Pagination(ctx echo.Context) error
}
