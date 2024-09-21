package contracts

import (
	"github.com/labstack/echo"
	"teukufuad/e-commerce/app/domains/models"
	"teukufuad/e-commerce/app/domains/types"
)

type TransactionRepository interface {
	Order(data models.Transaction) (*models.Transaction, error)
}

type TransactionServices interface {
	Order(request types.RequestOrder) (*types.ResponseOrder, error)
	//UpdateProfile() (*types.ResponseListCategory, error)
}

type TransactionController interface {
	Order(ctx echo.Context) error
	//UpdateProfile(ctx echo.Context) error
}
