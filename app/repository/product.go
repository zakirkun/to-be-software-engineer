package repository

import (
	"teukufuad/e-commerce/app/domains/contracts"
	"teukufuad/e-commerce/app/domains/models"
	"teukufuad/e-commerce/pkg/database"
)

type ProductRepository struct{}

func (p ProductRepository) Insert(data models.Product) (*models.Product, error) {
	db, err := database.DB.OpenDB()
	if err != nil {
		return nil, *err
	}

	if err := db.Create(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

func (p ProductRepository) GetAll() ([]*models.Product, error) {
	db, err := database.DB.OpenDB()
	if err != nil {
		return nil, *err
	}

	var data []*models.Product
	if err := db.Debug().Model(&models.Product{}).Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (p ProductRepository) Get(productId int) (*models.Product, error) {
	db, err := database.DB.OpenDB()
	if err != nil {
		return nil, *err
	}

	var data models.Product
	if err := db.Debug().Model(&models.Product{}).Where("id = ?", productId).Find(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

func (p ProductRepository) Update(productId int, data models.Product) (*models.Product, error) {
	db, err := database.DB.OpenDB()
	if err != nil {
		return nil, *err
	}

	if err := db.Debug().Model(&models.Product{}).Where("id = ?", productId).Save(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

func (p ProductRepository) Delete(productId int) error {
	db, err := database.DB.OpenDB()
	if err != nil {
		return *err
	}

	if err := db.Debug().Delete(&models.Product{}, productId).Error; err != nil {
		return err
	}

	return nil
}

func NewProductRepository() contracts.ProductRepository {
	return ProductRepository{}
}
