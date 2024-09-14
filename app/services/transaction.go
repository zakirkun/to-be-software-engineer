package services

import (
	"imzakir.dev/e-commerce/app/domains/contracts"
	"imzakir.dev/e-commerce/app/domains/models"
	"imzakir.dev/e-commerce/app/domains/types"
	"imzakir.dev/e-commerce/app/repository"
)

type transactionService struct{}

func (t transactionService) AddTransaction(request types.RequestCreateTransaction, username string) (*types.ResponseCreateTransaction, error) {
	repo := repository.NewTransactionRepository()
	repoCustomer := repository.NewCustomerRepository()

	customer, err := repoCustomer.GetByUsername(username)
	if err != nil {
		return nil, err
	}

	data, err := repo.Insert(models.Transaction{
		IdCustomer: customer.Id,
		IdProduct:  request.IdProduct,
		Qty:        request.Qty,
		Amount:     request.Amount,
	})

	if err != nil {
		return nil, err
	}

	return &types.ResponseCreateTransaction{
		Transaction: data,
	}, nil
}

func NewTransactionService() contracts.TransactionServices {
	return transactionService{}
}
