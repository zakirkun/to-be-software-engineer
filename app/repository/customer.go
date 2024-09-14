package repository

import (
	"imzakir.dev/e-commerce/app/domains/contracts"
	"imzakir.dev/e-commerce/app/domains/models"
	"imzakir.dev/e-commerce/pkg/database"
)

type CustomerRepository struct {
}

func (c CustomerRepository) GetByEmail(email string) (*models.Customer, error) {
	db, err := database.DB.OpenDB()
	if err != nil {
		return nil, *err
	}

	var customer models.Customer
	if err := db.Debug().Model(&models.Customer{}).Where("email = ?", email).Find(&customer).Error; err != nil {
		return nil, err
	}

	return &customer, nil
}

func (c CustomerRepository) Insert(data models.Customer) (*models.Customer, error) {
	db, err := database.DB.OpenDB()
	if err != nil {
		return nil, *err
	}

	if err := db.Create(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

func (c CustomerRepository) Update(id int, data models.Customer) (*models.Customer, error) {
	db, err := database.DB.OpenDB()
	if err != nil {
		return nil, *err
	}

	if err := db.Debug().Model(&models.Category{}).Where("id = ?", id).Save(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

func NewCustomerRepository() contracts.CustomerRepository {
	return CustomerRepository{}
}
