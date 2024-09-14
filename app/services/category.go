package services

import (
	"imzakir.dev/e-commerce/app/domains/contracts"
	"imzakir.dev/e-commerce/app/domains/models"
	"imzakir.dev/e-commerce/app/domains/types"
	"imzakir.dev/e-commerce/app/repository"
)


func NewCategoryServices() contracts.CategoryServices {
	return categoryServices{}
}

type categoryServices struct{}

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



// GetId implements contracts.CategoryServices.
func (c categoryServices) GetCategoryById(id int) (*types.ResponseCreateCategory, error) {
	data,err := repository.NewCategoryRepository().FindCategoryById(id)
	if err != nil {
		return nil, err
	}

	return &types.ResponseCreateCategory{
		Category: data,
	}, nil
}


//Delete By Id
func (c categoryServices) DeleteId(id int) (*types.ResponseCreateCategory, error) {
	data, err := repository.NewCategoryRepository().DeleteById(id)
	if err != nil {
		return nil,err
	}

	return  &types.ResponseCreateCategory{
		Category: data,
	}, nil
}

func (c categoryServices) UpdateCategoryById(id int, request types.RequestCreateCategory) (*types.ResponseCreateCategory, error){

	repo := repository.NewCategoryRepository()

	data, err := repo.Update(id, models.Category{
		CategoryName: request.Name,
	})

	if err != nil {
		return nil, err
	}

	return &types.ResponseCreateCategory{
		Category: data,
	}, nil

}