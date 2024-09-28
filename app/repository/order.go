package repository

import (
	"imzakir.dev/e-commerce/app/domains/contracts"
	"imzakir.dev/e-commerce/app/domains/models"
	"imzakir.dev/e-commerce/pkg/database"
)

type orderRepository struct{}

// Create implements contracts.OrderRepository.
func (o orderRepository) Create(data *models.Transaction) error {
	db, err := database.DB.OpenDB()
	if err != nil {
		return *err
	}

	// Note the use of tx as the database handle once you are within a transaction
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Create(&data).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// GetByID implements contracts.OrderRepository.
func (o orderRepository) GetByID(id uint) (*models.Transaction, error) {
	db, err := database.DB.OpenDB()
	if err != nil {
		return nil, *err
	}

	var data models.Transaction
	if err := db.Debug().Preload("Product").Preload("Customer").Model(&models.Transaction{}).Where("id = ?", id).Find(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

func NewOrderRepository() contracts.OrderRepository {
	return orderRepository{}
}
