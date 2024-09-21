package repository

import (
	"log"

	"imzakir.dev/e-commerce/app/domains/contracts"
	"imzakir.dev/e-commerce/app/domains/models"
	"imzakir.dev/e-commerce/pkg/database"
)

type categoryRepository struct{}

// GetByID implements contracts.CategoryRepository.
func (c categoryRepository) GetByID(id uint) (*models.Category, error) {
	db, err := database.DB.OpenDB()
	if err != nil {
		return nil, *err
	}

	var data models.Category
	if err := db.Model(&models.Category{}).Where("id = ?", id).Find(&data).Error; err != nil {
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

func (c categoryRepository) FindCategoryById(id int) (*models.Category, error){
	db, err := database.DB.OpenDB()
	if err != nil {
		return nil, *err
	}

	var data models.Category
	if err := db.Debug().Model(&models.Category{}).Where("id = ?",id).First(&data).Error; err != nil {
		return nil, err
	}
	return &data , nil
}

func (c categoryRepository) DeleteById(id int) (*models.Category, error) {
	db, err := database.DB.OpenDB()

	if err != nil {
		return nil, *err
	}

	var data models.Category
	if err := db.Debug().Model(&models.Category{}).Where("id = ?", id).Delete(&data).Error; err != nil {
		return nil,err
	}
	log.Println("error delete :",err)
	return  &data, nil

}

func (c categoryRepository) Update(id int, data models.Category) (*models.Category, error){
	db, err := database.DB.OpenDB()

	if err != nil {
		return nil, *err
		
	}

	if err := db.Debug().Model(&models.Category{}).Where("id = ?", id).Updates(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}


func NewCategoryRepository() contracts.CategoryRepository {
	return categoryRepository{}
}