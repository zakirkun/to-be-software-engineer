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

// Delete implements contracts.CategoryServices.
func (c categoryServices) Delete(id int) (bool, error) {
	repo := repository.NewCategoryRepository()
	dataCategory, err := repo.GetDetail(id)
	if err != nil {
		return false, err
	}

	data, err := repo.Delete(*dataCategory)

	if err != nil {
		return false, err
	}
	return data, err
}

// Update implements contracts.CategoryServices.
func (c categoryServices) Update(id int, request types.RequestCreateCategory) (*types.ResponseCreateCategory, error) {
	repo := repository.NewCategoryRepository()
	dataCategory, err := repo.GetDetail(id)
	if err != nil {
		return nil, err
	}

	dataCategory.CategoryName = request.Name

	data, err := repo.Update(*dataCategory)

	//data, err := repo.Update(*dataCategory)

	if err != nil {
		return nil, err
	}

	return &types.ResponseCreateCategory{
		Category: data,
	}, nil
}

// GetDetail implements contracts.CategoryServices.
func (c categoryServices) GetDetail(id int) (*types.ResponseGetDetailCategory, error) {
	repo := repository.NewCategoryRepository()

	// check cache first
	_id := strconv.Itoa(int(id))
	getKey, err := cache.CACHE.Get(context.Background(), fmt.Sprintf("category:%v", _id)).Result()
	if err == redis.Nil {
		getCategory, err := repo.GetDetail(id)
		if err != nil {
			return nil, err
		}

		if getCategory.Id == 0 {
			return nil, errors.New("record not found")
		}

		toJson := utils.StructToJson(&getCategory)

		cache.CACHE.Set(context.Background(), fmt.Sprintf("category:%v", _id), toJson, time.Duration(time.Minute*30))
		return &types.ResponseGetDetailCategory{
			Category: getCategory,
		}, nil
	}

	var parse models.Category
	if ok := utils.JsonToSruct([]byte(getKey), &parse); !ok {
		return nil, errors.ErrUnsupported
	}

	return &types.ResponseGetDetailCategory{
		Category: &parse,
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
