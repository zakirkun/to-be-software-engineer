package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"imzakir.dev/e-commerce/app/domains/contracts"
	"imzakir.dev/e-commerce/app/domains/types"
	"imzakir.dev/e-commerce/app/services"
	"imzakir.dev/e-commerce/utils"
)

type productController struct{}

// Delete implements contracts.ProductController.
func (p productController) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	if id == "" {
		return ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse(http.StatusBadRequest, "ERROR_VALIDATION", errors.New("invalid id")))
	}

	_id, _ := strconv.Atoi(id)
	svc := services.NewProductServices()
	if err := svc.Delete(uint(_id)); err != nil {
		return ctx.JSON(http.StatusInternalServerError, utils.SetErrorResponse(http.StatusInternalServerError, "INTERNAL_ERROR", err))
	}

	return ctx.JSON(http.StatusOK, utils.SetSuccessReponse(http.StatusOK, "SUCCESS", nil))
}

// GetAll implements contracts.ProductController.
func (p productController) GetAll(ctx echo.Context) error {
	svc := services.NewProductServices()
	data, err := svc.GetAll()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, utils.SetErrorResponse(http.StatusInternalServerError, "INTERNAL_ERROR", err))
	}

	return ctx.JSON(http.StatusOK, utils.SetSuccessReponse(http.StatusOK, "SUCCESS", data))
}

// GetByID implements contracts.ProductController.
func (p productController) GetByID(ctx echo.Context) error {
	id := ctx.Param("id")
	_id, _ := strconv.Atoi(id)
	svc := services.NewProductServices()

	data, err := svc.GetByID(uint(_id))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, utils.SetErrorResponse(http.StatusInternalServerError, "INTERNAL_ERROR", err))
	}

	return ctx.JSON(http.StatusOK, utils.SetSuccessReponse(http.StatusOK, "SUCCESS", data))
}

// Insert implements contracts.ProductController.
func (p productController) Insert(ctx echo.Context) error {
	var request types.RequestCreateProduct
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse(http.StatusBadRequest, "ERROR_VALIDATION", err))
	}

	svc := services.NewProductServices()
	data, err := svc.Create(request)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, utils.SetErrorResponse(http.StatusInternalServerError, "INTERNAL_ERROR", err))
	}

	return ctx.JSON(http.StatusCreated, utils.SetSuccessReponse(http.StatusCreated, "SUCCESS", data))
}

// Pagination implements contracts.ProductController.
func (p productController) Pagination(ctx echo.Context) error {
	data, err := services.NewProductServices().Pagination(ctx)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, utils.SetErrorResponse(http.StatusInternalServerError, "INTERNAL_ERROR", err))
	}

	return ctx.JSON(http.StatusOK, utils.SetSuccessReponse(http.StatusOK, "SUCCESS", data))
}

// Update implements contracts.ProductController.
func (p productController) Update(ctx echo.Context) error {
	var reuqest types.RequestCreateProduct
	if err := ctx.Bind(&reuqest); err != nil {
		return ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse(http.StatusBadRequest, "ERROR_VALIDATION", err))
	}

	id := ctx.Param("id")
	if id == "" {
		return ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse(http.StatusBadRequest, "ERROR_VALIDATION", errors.New("invalid id")))
	}

	_id, _ := strconv.Atoi(id)
	svc := services.NewProductServices()
	updated, err := svc.Update(uint(_id), reuqest)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, utils.SetErrorResponse(http.StatusInternalServerError, "INTERNAL_ERROR", err))
	}

	return ctx.JSON(http.StatusOK, utils.SetSuccessReponse(http.StatusOK, "SUCCESS", updated))
}

func NewProductController() contracts.ProductController {
	return productController{}
}
