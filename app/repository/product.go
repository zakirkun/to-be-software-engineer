package repository

import (
	"imzakir.dev/e-commerce/app/domains/contracts"
	"imzakir.dev/e-commerce/app/domains/models"
	"imzakir.dev/e-commerce/pkg/database"
)

type productRepository struct{}

// Create implements contracts.ProductRepository.
func (p productRepository) Create(data models.Product) error {
	db, err := database.DB.OpenDB()
	if err != nil {
		return *err
	}

	if err := db.Create(&data).Error; err != nil {
		return err
	}

	return nil
}

// Delete implements contracts.ProductRepository.
func (p productRepository) Delete(id uint) error {
	db, err := database.DB.OpenDB()
	if err != nil {
		return *err
	}

	if err := db.Where("id = ?", id).Delete(&models.Product{}).Error; err != nil {
		return err
	}

	return nil
}

// FindBy implements contracts.ProductRepository.
func (p productRepository) FindBy(where map[string]interface{}) (*[]models.Product, error) {
	db, err := database.DB.OpenDB()
	if err != nil {
		return nil, *err
	}

	var data []models.Product
	if err := db.Debug().Model(&models.Product{}).Where(where).Find(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

// GetAll implements contracts.ProductRepository.
func (p productRepository) GetAll() (*[]models.Product, error) {
	db, err := database.DB.OpenDB()
	if err != nil {
		return nil, *err
	}

	var data []models.Product
	if err := db.Model(&models.Product{}).Find(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

// Update implements contracts.ProductRepository.
func (p productRepository) Update(id uint, data models.Product) error {
	db, err := database.DB.OpenDB()
	if err != nil {
		return *err
	}

	if err := db.Where("id = ?", id).Updates(&data).Error; err != nil {
		return err
	}

	return nil
}

func NewProductRepository() contracts.ProductRepository {
	return productRepository{}
}
