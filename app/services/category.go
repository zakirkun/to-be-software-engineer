package services

import (
	"errors"

	"imzakir.dev/e-commerce/app/domains/contracts"
	"imzakir.dev/e-commerce/app/domains/models"
	"imzakir.dev/e-commerce/app/domains/types"
	"imzakir.dev/e-commerce/app/repository"
)

type categoryServices struct{}

// GetDetail implements contracts.CategoryServices.
func (c categoryServices) GetDetail(id int) (*types.ResponseGetDetailCategory, error) {
	repo := repository.NewCategoryRepository()
	data, err := repo.GetDetail(id)
	if err != nil {
		return nil, err
	}

	if data.Id == 0 {
		return nil, errors.New("Record Not Found")
	}

	return &types.ResponseGetDetailCategory{
		Category: data,
	}, nil
}

// GetAll implements contracts.CategoryServices.
func (c categoryServices) GetAll() (*types.ResponseListCategory, error) {
	repo := repository.NewCategoryRepository()
	data, err := repo.GetAll()
	if err != nil {
		return nil, err
	}

	return &types.ResponseListCategory{
		Category: data,
	}, nil
}

// Insert implements contracts.CategoryServices.
func (c categoryServices) Insert(request types.RequestCreateCategory) (*types.ResponseCreateCategory, error) {
	repo := repository.NewCategoryRepository()
	data, err := repo.Insert(models.Category{
		CategoryName: request.Name,
	})
	if err != nil {
		return nil, err
	}

	return &types.ResponseCreateCategory{
		Category: data,
	}, nil
}

func NewCategoryServices() contracts.CategoryServices {
	return categoryServices{}
}
