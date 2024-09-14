package repository

import (
	"imzakir.dev/e-commerce/app/domains/contracts"
	"imzakir.dev/e-commerce/app/domains/models"
	"imzakir.dev/e-commerce/pkg/database"
)

type transactionRepository struct {
}

func (t transactionRepository) Insert(data models.Transaction) (*models.Transaction, error) {
	db, err := database.DB.OpenDB()
	if err != nil {
		return nil, *err
	}

	if err := db.Create(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

func NewTransactionRepository() contracts.TransactionRepository {
	return transactionRepository{}
}
