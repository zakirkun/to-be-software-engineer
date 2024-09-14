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


func NewCategoryServices() contracts.CategoryServices {
	return categoryServices{}
}

type categoryServices struct{}

// GetByID implements contracts.CategoryServices.
func (c categoryServices) GetByID(id uint) (*types.ResponseCreateCategory, error) {

	repo := repository.NewCategoryRepository()

	// check cache first
	_id := strconv.Itoa(int(id))
	getKey, err := cache.CACHE.Get(context.Background(), fmt.Sprintf("category:%v", _id)).Result()
	if err == redis.Nil {
		getCategory, err := repo.GetByID(id)
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