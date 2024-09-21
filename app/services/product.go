package services

import (
	"imzakir.dev/e-commerce/app/domains/contracts"
	"imzakir.dev/e-commerce/app/domains/models"
	"imzakir.dev/e-commerce/app/domains/types"
	"imzakir.dev/e-commerce/app/repository"
)

func NewProductServices() contracts.ProductServices {
	return productServices{}
}

type productServices struct{}

// GetAll implements contracts.CategoryServices.
func (c productServices) GetAll() (*types.ResponseListProduct, error) {
	repo := repository.NewProductRepository()
	data, err := repo.GetAll()
	if err != nil {
		return nil, err
	}

	return &types.ResponseListProduct{
		Product: data,
	}, nil
}

// Insert implements contracts.CategoryServices.
func (c productServices) Insert(request types.RequestCreateProduct) (*types.ResponseCreateProduct, error) {
	repo := repository.NewProductRepository()
	data, err := repo.Insert(models.Product{
		ProductName: request.NameProduct,
		ProductImage: request.NameImage,
		ProductDescription: request.NameProduction,
		Price: request.NamePrice,
		CategoryId: request.NameCategoryId,

	})


	if err != nil {
		return nil, err
	}

	return &types.ResponseCreateProduct{
		Product: data,
	}, nil
}



// GetId implements contracts.CategoryServices.
func (c productServices) GetProductById(id int) (*types.ResponseCreateProduct, error) {
	data,err := repository.NewProductRepository().FindProductById(id)
	if err != nil {
		return nil, err
	}

	return &types.ResponseCreateProduct{
		Product: data,
	}, nil
}


//Delete By Id
func (c productServices) DeleteId(id int) (*types.ResponseCreateProduct, error) {
	data, err := repository.NewProductRepository().DeleteById(id)
	if err != nil {
		return nil,err
	}

	return  &types.ResponseCreateProduct{
		Product: data,
	}, nil
}


func (c productServices) Update(id int, request types.RequestCreateProduct) (*types.ResponseCreateProduct, error) {
	repo := repository.NewProductRepository()
	data, err := repo.Update(id, models.Product{
		ProductName: request.NameProduct,
		ProductImage: request.NameImage,
		ProductDescription: request.NameProduction,
		Price: request.NamePrice,
		CategoryId: request.NameCategoryId,

	})


	if err != nil {
		return nil, err
	}

	return &types.ResponseCreateProduct{
		Product: data,
	}, nil
}
