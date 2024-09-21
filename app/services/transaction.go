package services

import (
	"imzakir.dev/e-commerce/app/domains/contracts"
	"imzakir.dev/e-commerce/app/domains/models"
	"imzakir.dev/e-commerce/app/domains/types"
	"imzakir.dev/e-commerce/app/repository"

)


func NewTransactionServices() contracts.TransactionServices {
	return transactionServices{}
}

type transactionServices struct{}




// Insert implements contracts.CategoryServices.
func (c transactionServices) Insert(request types.RequestCreateTransaction) (*types.ResponseCreateTransaction, error) {
	repo := repository.NewTransactionRepository()
	data, err := repo.Insert(models.Transaction{
		ProductId: request.ProductId,
		CustomerId: request.CustomerId,
		Qty : request.Qty,
		Amount : request.Amount,
	})
	if err != nil {
		return nil, err
	}

	return &types.ResponseCreateTransaction{
		Transaction: data,
	}, nil
}



// GetId implements contracts.CategoryServices.
func (c transactionServices) GetTransactionById(id int) (*types.ResponseCreateTransaction, error) {
	data,err := repository.NewTransactionRepository().FindTransactionById(id)
	if err != nil {
		return nil, err
	}

	return &types.ResponseCreateTransaction{
		Transaction: data,
	}, nil
}

