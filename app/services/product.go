package services

import (
	"imzakir.dev/e-commerce/app/domains/contracts"
	"imzakir.dev/e-commerce/app/domains/models"
	"imzakir.dev/e-commerce/app/domains/types"
	"imzakir.dev/e-commerce/app/repository"
)

type productService struct {
}

func (p productService) GetDetail(productId int) (*models.Product, error) {
	repo := repository.NewProductRepository()
	data, err := repo.FindById(productId)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (p productService) GetAllProducts() (*[]models.Product, error) {
	repo := repository.NewProductRepository()
	data, err := repo.FindAll()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (p productService) AddProduct(request types.RequestCreateProduct) (*types.ResponseCreateProduct, error) {
	repo := repository.NewProductRepository()
	data, err := repo.Insert(models.Product{
		CategoryId:         request.CategoryId,
		ProductName:        request.ProductName,
		ProductDescription: request.ProductDescription,
		ProductImage:       request.ProductImage,
		Price:              request.Price,
	})

	if err != nil {
		return nil, err
	}

	return &types.ResponseCreateProduct{
		Product: data,
	}, nil
}

func NewProductService() contracts.ProductServices {
	return productService{}
}
