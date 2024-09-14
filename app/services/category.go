package services

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"imzakir.dev/e-commerce/app/domains/contracts"
	"imzakir.dev/e-commerce/app/domains/models"
	"imzakir.dev/e-commerce/app/domains/types"
	"imzakir.dev/e-commerce/app/repository"
	"imzakir.dev/e-commerce/pkg/cache"
	"imzakir.dev/e-commerce/utils"
)

type categoryServices struct{}

func (c categoryServices) Delete(categoryId int) (bool, error) {
	repo := repository.NewCategoryRepository()
	dataCategory, err := repo.Show(categoryId)
	if err != nil {
		return false, err
	}

	data, err := repo.Delete(*dataCategory)

	if err != nil {
		return false, err
	}
	return data, err
}

func (c categoryServices) Update(categoryId int, request types.RequestCreateCategory) (*types.ResponseCreateCategory, error) {
	repo := repository.NewCategoryRepository()
	dataCategory, err := repo.Show(categoryId)
	if err != nil {
		return nil, err
	}

	dataCategory.CategoryName = request.CategoryName

	data, err := repo.Update(*dataCategory)

	//data, err := repo.Update(*dataCategory)

	if err != nil {
		return nil, err
	}

	return &types.ResponseCreateCategory{
		Category: data,
	}, nil
}

func (c categoryServices) Show(categoryId int) (*types.ResponseCreateCategory, error) {
	repo := repository.NewCategoryRepository()
	data, err := repo.Show(categoryId)

	if err != nil {
		return nil, err
	}

	return &types.ResponseCreateCategory{
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
		CategoryName: request.CategoryName,
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
