package services

import (
	"errors"
	"imzakir.dev/e-commerce/app/domains/contracts"
	"imzakir.dev/e-commerce/app/domains/models"
	"imzakir.dev/e-commerce/app/domains/types"
	"imzakir.dev/e-commerce/app/repository"
	"imzakir.dev/e-commerce/pkg/cache"
	"imzakir.dev/e-commerce/utils"
)

type categoryServices struct{}

func (c categoryServices) Get(categoryId int) (*types.ResponseCategory, error) {
	repo := repository.NewCategoryRepository()
	data, err := repo.Get(categoryId)
	if err != nil {
		return nil, err
	}

	if data.Id == 0 {
		return nil, errors.New("category not found")
	}

	return &types.ResponseCategory{
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

func (c categoryServices) Update(request types.RequestCreateCategory, categoryId int) (*types.ResponseCreateCategory, error) {
	repo := repository.NewCategoryRepository()

	category, err := repo.Get(categoryId)
	if err != nil {
		return nil, err
	}

	data, err := repo.Update(categoryId, category.ToUpdateCategory(request.Name))
	if err != nil {
		return nil, err
	}

	return &types.ResponseCreateCategory{
		Category: data,
	}, nil
}

func (c categoryServices) Delete(categoryId int) error {
	repo := repository.NewCategoryRepository()

	err := repo.Delete(categoryId)
	if err != nil {
		return err
	}

	return nil
}

func NewCategoryServices() contracts.CategoryServices {
	return categoryServices{}
}
