package contracts

import (
	"github.com/labstack/echo"
	"imzakir.dev/e-commerce/app/domains/models"
	"imzakir.dev/e-commerce/app/domains/types"
)

type CategoryRepository interface {
	Insert(data models.Category) (*models.Category, error)
	GetAll() (*[]models.Category, error)
	FindCategoryById(id int) (*models.Category, error)
	Update(id int, data models.Category) (*models.Category, error)
	DeleteById(id int ) (*models.Category,error)
}

type CategoryServices interface {
	Insert(request types.RequestCreateCategory) (*types.ResponseCreateCategory, error)
	GetAll() (*types.ResponseListCategory, error)
	GetCategoryById(id int)  (*types.ResponseCreateCategory, error)
	UpdateCategoryById(id int, request types.RequestCreateCategory) (*types.ResponseCreateCategory, error)
	DeleteId (id int) (*types.ResponseCreateCategory, error)
}

type CategoryController interface {
	Insert(ctx echo.Context) error
	GetAll(ctx echo.Context) error
	Show(ctx echo.Context) error
	Edit(ctx echo.Context) error
	Delete(ctx echo.Context) error
} 
