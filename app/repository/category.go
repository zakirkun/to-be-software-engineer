package repository

import (
	"teukufuad/e-commerce/app/domains/contracts"
	"teukufuad/e-commerce/app/domains/models"
	"teukufuad/e-commerce/pkg/database"
)

type categoryRepository struct{}

func (c categoryRepository) Get(categoryId int) (*models.Category, error) {
	db, err := database.DB.OpenDB()
	if err != nil {
		return nil, *err
	}

	var data models.Category
	if err := db.Debug().Model(&models.Category{}).Where("id = ?", categoryId).Find(&data).Error; err != nil {
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

func (c categoryRepository) Update(categoryId int, data models.Category) (*models.Category, error) {
	db, err := database.DB.OpenDB()
	if err != nil {
		return nil, *err
	}

	if err := db.Debug().Model(&models.Category{}).Where("id = ?", categoryId).Save(&data).Error; err != nil {
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

func (c categoryRepository) Delete(categoryId int) error {
	db, err := database.DB.OpenDB()
	if err != nil {
		return *err
	}

	if err := db.Debug().Delete(&models.Category{}, categoryId).Error; err != nil {
		return err
	}

	return nil
}

func NewCategoryRepository() contracts.CategoryRepository {
	return categoryRepository{}
}
