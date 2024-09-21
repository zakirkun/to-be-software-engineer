package repository

import (
	"imzakir.dev/e-commerce/app/domains/contracts"
	"imzakir.dev/e-commerce/app/domains/models"
	"imzakir.dev/e-commerce/pkg/database"
)

type categoryRepository struct{}

// GetByCategory implements contracts.CategoryRepository.
func (c categoryRepository) GetByCategory(cat_id uint) (*models.Category, error) {
	db, err := database.DB.OpenDB()
	if err != nil {
		return nil, *err
	}

	var data models.Category
	if err := db.Preload("Products").Where("id = ?", cat_id).Find(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

// Delete im	plements contracts.CategoryRepository.
func (c categoryRepository) Delete(data models.Category) (bool, error) {
	db, err := database.DB.OpenDB()
	if err != nil {
		return false, *err
	}

	if err := db.Delete(&data).Error; err != nil {
		return false, nil
	}

	return true, nil

}

// Update implements contracts.CategoryRepository.
func (c categoryRepository) Update(data models.Category) (*models.Category, error) {
	db, err := database.DB.OpenDB()
	if err != nil {
		return nil, *err
	}

	if err := db.Save(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

// GetDetail implements contracts.CategoryRepository.
func (c categoryRepository) GetDetail(id int) (*models.Category, error) {
	db, err := database.DB.OpenDB()
	if err != nil {
		return nil, *err
	}

	var data models.Category
	if err := db.Debug().Model(&models.Category{}).Where("id = ?", id).Find(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

// GetAll implements contracts.CategoryRepository.
func (c categoryRepository) GetAll() (*[]models.Category, error) {
	db, err := database.DB.OpenDB()
	if err != nil {
		return nil, *err
	}

	var data []models.Category
	if err := db.Debug().Model(&models.Category{}).Find(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

func NewCategoryRepository() contracts.CategoryRepository {
	return categoryRepository{}
}

// Insert implements contracts.CategoryRepository.
func (c categoryRepository) Insert(data models.Category) (*models.Category, error) {
	db, err := database.DB.OpenDB()
	if err != nil {
		return nil, *err
	}

	if err := db.Create(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

// GetDetail implements contracts.CategoryRepository
