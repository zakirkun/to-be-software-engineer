package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"imzakir.dev/e-commerce/app/domains/contracts"
	"imzakir.dev/e-commerce/app/domains/models"
	"imzakir.dev/e-commerce/app/domains/types"
	"imzakir.dev/e-commerce/app/repository"
	"imzakir.dev/e-commerce/pkg/cache"
	"imzakir.dev/e-commerce/utils"
	"strconv"
	"time"
)

type categoryServices struct{}

func (c categoryServices) Delete(categoryId int) (bool, error) {
	repo := repository.NewCategoryRepository()
	getCategory, err := repo.Show(categoryId)
	if err != nil {
		return false, nil
	}

	if getCategory.Id == 0 {
		return false, errors.New("record not found")
	}

	data, err := repo.Delete(*getCategory)

	_id := strconv.Itoa(getCategory.Id)

	// Delete cache
	cache.CACHE.Del(context.Background(), fmt.Sprintf("category:%v", _id))

	return data, nil
}

func (c categoryServices) Update(categoryId int, request types.RequestCreateCategory) (*types.ResponseCreateCategory, error) {
	repo := repository.NewCategoryRepository()

	getCategory, err := repo.Show(categoryId)
	if err != nil {
		return nil, err
	}

	if getCategory.Id == 0 {
		return nil, errors.New("record not found")
	}

	getCategory.CategoryName = request.CategoryName

	data, err := repo.Update(*getCategory)

	if err != nil {
		return nil, err
	}
	_id := strconv.Itoa(getCategory.Id)

	// Delete cache
	cache.CACHE.Del(context.Background(), fmt.Sprintf("category:%v", _id))

	return &types.ResponseCreateCategory{
		Category: data,
	}, nil
}

func (c categoryServices) Show(categoryId int) (*types.ResponseCreateCategory, error) {
	repo := repository.NewCategoryRepository()

	// check cache first
	_id := strconv.Itoa(categoryId)
	getKey, err := cache.CACHE.Get(context.Background(), fmt.Sprintf("category:%v", _id)).Result()
	if err == redis.Nil {
		getCategory, err := repo.Show(categoryId)
		if err != nil {
			return nil, err
		}

		if getCategory.Id == 0 {
			return nil, errors.New("record not found")
		}

		toJson := utils.StructToJson(&getCategory)

		cache.CACHE.Set(context.Background(), fmt.Sprintf("category:%v", _id), toJson, time.Duration(time.Minute*30))
		return &types.ResponseCreateCategory{
			Category: getCategory,
		}, nil
	}

	var parse models.Category
	if ok := utils.JsonToSruct([]byte(getKey), &parse); !ok {
		return nil, errors.ErrUnsupported
	}

	return &types.ResponseCreateCategory{
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
