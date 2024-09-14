package contracts

import (
	"github.com/labstack/echo"
	"imzakir.dev/e-commerce/app/domains/models"
	"imzakir.dev/e-commerce/app/domains/types"
)

type TransactionRepository interface {
	Insert(data models.Transaction) (*models.Transaction, error)
}

type TransactionServices interface {
	AddTransaction(request types.RequestCreateTransaction, username string) (*types.ResponseCreateTransaction, error)
}

type TransactionController interface {
	Create(ctx echo.Context) error
}
