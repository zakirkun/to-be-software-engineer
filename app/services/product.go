package services

import (
	"github.com/labstack/echo"
	"github.com/morkid/paginate"
	"imzakir.dev/e-commerce/app/domains/contracts"
	"imzakir.dev/e-commerce/app/domains/models"
	"imzakir.dev/e-commerce/app/domains/types"
	"imzakir.dev/e-commerce/app/repository"
	"imzakir.dev/e-commerce/pkg/database"
)

type productServices struct{}

// Pagination implements contracts.ProductServices.
func (p productServices) Pagination(ctx echo.Context) (*paginate.Page, error) {
	db, err := database.DB.OpenDB()
	if err != nil {
		return nil, *err
	}

	stmt := db.Model(&models.Product{})
	pg := paginate.New()
	page := pg.With(stmt).Request(ctx.Request()).Response(&models.Product{})

	return &page, nil

}

// Create implements contracts.ProductServices.
func (p productServices) Create(request types.RequestCreateProduct) (*types.ResponseCreateProduct, error) {

	data := models.Product{
		CategoryID:         request.CategoryID,
		ProductName:        request.ProductName,
		ProductImage:       request.ProductImage,
		ProductDescription: request.ProductDescription,
		Price:              request.Price,
	}

	if err := repository.NewProductRepository().Create(data); err != nil {
		return nil, err
	}

	return &types.ResponseCreateProduct{
		Product: &data,
	}, nil
}

// Delete implements contracts.ProductServices.
func (p productServices) Delete(id uint) error {
	if err := repository.NewProductRepository().Delete(id); err != nil {
		return err
	}

	return nil
}

// GetAll implements contracts.ProductServices.
func (p productServices) GetAll() (*types.ResponsegetAllProduct, error) {
	getProduct, err := repository.NewProductRepository().GetAll()
	if err != nil {
		return nil, err
	}

	return &types.ResponsegetAllProduct{
		Product: getProduct,
	}, nil
}

// GetByID implements contracts.ProductServices.
func (p productServices) GetByID(id uint) (*types.ResponsegetAllProduct, error) {
	where := make(map[string]interface{})
	where["id"] = id

	getProduct, err := repository.NewProductRepository().FindBy(where)
	if err != nil {
		return nil, err
	}

	return &types.ResponsegetAllProduct{
		Product: getProduct,
	}, nil
}

// Update implements contracts.ProductServices.
func (p productServices) Update(id uint, request types.RequestCreateProduct) (*types.ResponseCreateProduct, error) {

	data := models.Product{
		CategoryID:         request.CategoryID,
		ProductName:        request.ProductName,
		ProductImage:       request.ProductImage,
		ProductDescription: request.ProductDescription,
		Price:              request.Price,
	}

	if err := repository.NewProductRepository().Update(id, data); err != nil {
		return nil, err
	}

	return &types.ResponseCreateProduct{
		Product: &data,
	}, nil
}

func NewProductServices() contracts.ProductServices {
	return productServices{}
}
