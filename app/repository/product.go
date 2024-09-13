package repository

import (
	"imzakir.dev/e-commerce/app/domains/contracts"
	"imzakir.dev/e-commerce/app/domains/models"
	"imzakir.dev/e-commerce/pkg/database"
)

type productRepository struct {
}

func (p productRepository) FindAll() (*[]models.Product, error) {
	db, err := database.DB.OpenDB()
	if err != nil {
		return nil, *err
	}

	var data []models.Product

	if err := db.Debug().Model(&models.Product{}).Find(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (p productRepository) Insert(data models.Product) (*models.Product, error) {
	db, err := database.DB.OpenDB()
	if err != nil {
		return nil, *err
	}

	if err := db.Create(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

func NewProductRepository() contracts.ProductRepository {
	return productRepository{}
}
