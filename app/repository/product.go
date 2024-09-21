package repository

import (
	"imzakir.dev/e-commerce/app/domains/contracts"
	"imzakir.dev/e-commerce/app/domains/models"
	"imzakir.dev/e-commerce/pkg/database"
)

type productRepository struct {
}

func (p productRepository) Delete(data models.Product) (bool, error) {
	db, err := database.DB.OpenDB()
	if err != nil {
		return false, *err
	}

	if err := db.Delete(&data).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (p productRepository) Update(data models.Product) (*models.Product, error) {
	db, err := database.DB.OpenDB()
	if err != nil {
		return nil, *err
	}

	if err := db.Save(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (p productRepository) FindById(productId int) (*models.Product, error) {
	db, err := database.DB.OpenDB()
	if err != nil {
		return nil, *err
	}

	var data models.Product

	if err := db.Debug().Model(&models.Product{}).Where("id = ?", productId).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
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
