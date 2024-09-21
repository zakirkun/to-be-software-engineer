package repository

import (
	"imzakir.dev/e-commerce/app/domains/contracts"
	"imzakir.dev/e-commerce/app/domains/models"
	"imzakir.dev/e-commerce/pkg/database"
)

type customerRepository struct{}




// Insert implements contracts.CategoryRepository.
func (c customerRepository) Insert(data models.Customer) (*models.Customer, error) {
	db, err := database.DB.OpenDB()
	if err != nil {
		return nil, *err
	}

	if err := db.Create(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

func (c customerRepository) FindCustomerById(id int) (*models.Customer, error){
	db, err := database.DB.OpenDB()
	if err != nil {
		return nil, *err
	}

	var data models.Customer
	if err := db.Debug().Model(&models.Customer{}).Where("id = ?",id).First(&data).Error; err != nil {
		return nil, err
	}
	return &data , nil
}




func NewCustomerRepository() contracts.CustomerRepository {
	return customerRepository{}
}