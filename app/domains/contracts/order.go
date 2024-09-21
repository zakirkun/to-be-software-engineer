package contracts

import (
	"github.com/labstack/echo"
	"imzakir.dev/e-commerce/app/domains/models"
	"imzakir.dev/e-commerce/app/domains/types"
)

type OrderRepository interface {
	Create(data models.Transaction) error
	GetByID(id uint) (*models.Transaction, error)
}

type OrderServices interface {
	CreateTransaction(request types.RequestCreateTransaction) (*types.ResponseGetTransaction, error)
	GetTransaction(id uint) (*types.ResponseGetTransaction, error)
	HandleSentEmail(data []byte) error
}

type OrderController interface {
	CreateTransaction(ctx echo.Context) error
	GetTransaction(ctx echo.Context) error
}
