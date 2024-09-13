package services

import (
	"errors"
	"imzakir.dev/e-commerce/app/domains/contracts"
	"imzakir.dev/e-commerce/app/domains/models"
	"imzakir.dev/e-commerce/app/domains/types"
	"imzakir.dev/e-commerce/app/repository"
)

type ProductService struct {
}

func (p ProductService) Insert(request types.RequestProduct) (*types.ResponseCreateProduct, error) {
	repo := repository.NewProductRepository()
	data, err := repo.Insert(models.Product{
		CategoryId:  request.CategoryId,
		Name:        request.Name,
		Image:       request.Image,
		Description: request.Description,
		Price:       request.Price,
	})

	if err != nil {
		return nil, err
	}

	return &types.ResponseCreateProduct{
		Product: data,
	}, nil
}

func (p ProductService) GetAll() (*types.ResponseListProduct, error) {
	repo := repository.NewProductRepository()
	data, err := repo.GetAll()
	if err != nil {
		return nil, err
	}

	return &types.ResponseListProduct{
		Product: data,
	}, nil
}

func (p ProductService) Get(productId int) (*types.ResponseProduct, error) {
	repo := repository.NewProductRepository()
	data, err := repo.Get(productId)
	if err != nil {
		return nil, err
	}

	if data.Id == 0 {
		return nil, errors.New("product not found")
	}

	return &types.ResponseProduct{
		Product: data,
	}, nil
}

func (p ProductService) Update(request types.RequestProduct, productId int) (*types.ResponseCreateProduct, error) {
	repo := repository.NewProductRepository()
	data, err := repo.Get(productId)
	if err != nil {
		return nil, err
	}

	if data.Id == 0 {
		return nil, errors.New("category not found")
	}

	data, err = repo.Update(productId, data.ToUpdateProduct(models.Product{
		CategoryId:  request.CategoryId,
		Name:        request.Name,
		Image:       request.Image,
		Description: request.Description,
		Price:       request.Price,
	}))

	if err != nil {
		return nil, err
	}

	return &types.ResponseCreateProduct{
		Product: data,
	}, nil
}

func (p ProductService) Delete(productId int) error {
	repo := repository.NewProductRepository()
	err := repo.Delete(productId)

	if err != nil {
		return err
	}

	return nil
}

func NewProductService() contracts.ProductServices {
	return ProductService{}
}
