package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"teukufuad/e-commerce/app/domains/contracts"
	"teukufuad/e-commerce/app/domains/models"
	"teukufuad/e-commerce/app/domains/types"
	"teukufuad/e-commerce/app/repository"
	"teukufuad/e-commerce/pkg/cache"
	"teukufuad/e-commerce/utils"
	"time"
)

type categoryServices struct{}

func (c categoryServices) Get(categoryId int) (*types.ResponseCategory, error) {
	repo := repository.NewCategoryRepository()

	getKey, err := cache.CACHE.Get(context.Background(), fmt.Sprintf("category:%v", categoryId)).Result()

	if err == redis.Nil {
		getCategory, err := repo.Get(categoryId)
		if err != nil {
			return nil, err
		}

		if getCategory.Id == 0 {
			return nil, errors.New("record not found")
		}

		toJson := utils.StructToJson(&getCategory)

		cache.CACHE.Set(context.Background(), fmt.Sprintf("category:%v", categoryId), toJson, time.Duration(time.Minute*30))
		return &types.ResponseCategory{
			Category: getCategory,
		}, nil
	}

	var parse models.Category
	if ok := utils.JsonToSruct([]byte(getKey), &parse); !ok {
		return nil, errors.New("internal server error")
	}

	return &types.ResponseCategory{
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
