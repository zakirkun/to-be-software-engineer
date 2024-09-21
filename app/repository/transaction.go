package repository

import (
	"teukufuad/e-commerce/app/domains/contracts"
	"teukufuad/e-commerce/app/domains/models"
	"teukufuad/e-commerce/pkg/database"
)

type TransactionRepository struct {
}

func (t TransactionRepository) Order(data models.Transaction) (*models.Transaction, error) {
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
	return TransactionRepository{}
}
