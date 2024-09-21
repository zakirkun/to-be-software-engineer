package repository

import (
	"imzakir.dev/e-commerce/app/domains/contracts"
	"imzakir.dev/e-commerce/app/domains/models"
	"imzakir.dev/e-commerce/pkg/database"
)

type customerRepository struct{}

// Create implements contracts.CustomerRepository.
func (c customerRepository) Create(data models.Customer) error {
	db, err := database.DB.OpenDB()
	if err != nil {
		return *err
	}

	if err := db.Create(&data).Error; err != nil {
		return err
	}

	return nil
}

// GetWhere implements contracts.CustomerRepository.
func (c customerRepository) GetWhere(where map[string]interface{}) (*models.Customer, error) {
	db, err := database.DB.OpenDB()
	if err != nil {
		return nil, *err
	}

	var data models.Customer
	if err := db.Debug().Model(&models.Customer{}).
		Where(where).Find(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

// Update implements contracts.CustomerRepository.
func (c customerRepository) Update(id uint, data models.Customer) error {
	db, err := database.DB.OpenDB()
	if err != nil {
		return *err
	}

	if err := db.Debug().Model(&models.Customer{}).Where("id = ?", id).
		Updates(&data).Error; err != nil {
		return err
	}

	return nil
}

func NewCustomerRepository() contracts.CustomerRepository {
	return customerRepository{}
}
