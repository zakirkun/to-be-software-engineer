package contracts

import (
	"github.com/labstack/echo"
	"imzakir.dev/e-commerce/app/domains/models"
	"imzakir.dev/e-commerce/app/domains/types"
)

type TransactionRepository interface {
	Insert(data models.Transaction) (*models.Transaction, error)
	FindTransactionById(id int) (*models.Transaction, error)
	
}

type TransactionServices interface {
	Insert(request types.RequestCreateTransaction) (*types.ResponseCreateTransaction, error)
	GetTransactionById(id int)  (*types.ResponseCreateTransaction, error)

}

type TransactionController interface {
	Insert(ctx echo.Context) error
	Show(ctx echo.Context) error
}
