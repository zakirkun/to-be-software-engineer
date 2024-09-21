package services

import (
	"teukufuad/e-commerce/app/domains/contracts"
	"teukufuad/e-commerce/app/domains/models"
	"teukufuad/e-commerce/app/domains/types"
	"teukufuad/e-commerce/app/repository"
)

type TransactionService struct {
}

func (t TransactionService) Order(request types.RequestOrder) (*types.ResponseOrder, error) {
	repo := repository.NewTransactionRepository()
	//auth := jwt.NewJWTImpl(config.GetString("jwt.signature_key"), config.GetInt("jwt.day_expired"))
	//calims := auth.ParseToken()
	data, err := repo.Order(models.Transaction{
		ProductId:  request.ProductId,
		CustomerId: 0,
		Qty:        request.Qty,
		Amount:     request.Amount,
	})

	if err != nil {
		return nil, err
	}

	return &types.ResponseOrder{
		Transaction: data,
	}, nil
}

func NewTransactionService() contracts.TransactionServices {
	return TransactionService{}
}
