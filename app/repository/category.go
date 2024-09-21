package repository

import (
	"imzakir.dev/e-commerce/app/domains/contracts"
	"imzakir.dev/e-commerce/app/domains/models"
	"imzakir.dev/e-commerce/pkg/database"
)

type categoryRepository struct{}

func (c categoryRepository) Delete(data models.Category) (bool, error) {
	db, err := database.DB.OpenDB()
	if err != nil {
		return false, *err
	}

	if err := db.Delete(&data).Error; err != nil {
		return false, err
	}

	return true, nil
}

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

func (c categoryRepository) Show(categoryId int) (*models.Category, error) {
	db, err := database.DB.OpenDB()
	if err != nil {
		return nil, *err
	}

	var category models.Category
	if err := db.Debug().Model(&models.Category{}).Where("id = ?", categoryId).First(&category).Error; err != nil {
		return nil, err
	}

	return &category, nil
}

// GetAll implements contracts.CategoryRepository.
func (c categoryRepository) GetAll() (*[]models.Category, error) {
	db, err := database.DB.OpenDB()
	if err != nil {
		return nil, *err
	}

	var data []models.Category
	if err := db.Preload("Products").Debug().Model(&models.Category{}).Find(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
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

func NewCategoryRepository() contracts.CategoryRepository {
	return categoryRepository{}
}
