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

type categoryController struct{}

// Pagination implements contracts.CategoryController.
func (c categoryController) Pagination(ctx echo.Context) error {

	data, err := services.NewCategoryServices().Pagination(ctx)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, utils.SetErrorResponse(http.StatusInternalServerError, "INTERNAL_ERROR", err))
	}

	return ctx.JSON(http.StatusOK, utils.SetSuccessReponse(http.StatusOK, "SUCCESS", data))
}

// GetByCategory implements contracts.CategoryController.
func (c categoryController) GetByCategory(ctx echo.Context) error {
	id := ctx.Param("id")
	if id == "" {
		return ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse(http.StatusBadRequest, "ERROR_VALIDATION", errors.New("invalid id")))
	}

	_id, _ := strconv.Atoi(id)
	svc := services.NewCategoryServices()
	data, err := svc.GetByCategory(uint(_id))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, utils.SetErrorResponse(http.StatusInternalServerError, "INTERNAL_ERROR", err))
	}

	return ctx.JSON(http.StatusOK, utils.SetSuccessReponse(http.StatusOK, "SUCCESS", data))
}

// Delete implements contracts.CategoryController.
func (c categoryController) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	if id == "" {
		return ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse(http.StatusBadRequest, "ERROR_VALIDATION", errors.New("invalid id")))
	}

	_id, _ := strconv.Atoi(id)
	svc := services.NewCategoryServices()
	if err := svc.Delete(uint(_id)); err != nil {
		return ctx.JSON(http.StatusInternalServerError, utils.SetErrorResponse(http.StatusInternalServerError, "INTERNAL_ERROR", err))
	}

	return ctx.JSON(http.StatusOK, utils.SetSuccessReponse(http.StatusOK, "SUCCESS", nil))
}

// Update implements contracts.CategoryController.
func (c categoryController) Update(ctx echo.Context) error {

	var reuqest types.RequestUpdateCategory
	if err := ctx.Bind(&reuqest); err != nil {
		return ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse(http.StatusBadRequest, "ERROR_VALIDATION", err))
	}

	id := ctx.Param("id")
	if id == "" {
		return ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse(http.StatusBadRequest, "ERROR_VALIDATION", errors.New("invalid id")))
	}

	_id, _ := strconv.Atoi(id)
	svc := services.NewCategoryServices()
	updated, err := svc.Update(uint(_id), reuqest)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, utils.SetErrorResponse(http.StatusInternalServerError, "INTERNAL_ERROR", err))
	}

	return ctx.JSON(http.StatusOK, utils.SetSuccessReponse(http.StatusOK, "SUCCESS", updated))
}

// GetByID implements contracts.CategoryController.
func (c categoryController) GetByID(ctx echo.Context) error {

	id := ctx.Param("id")
	_id, _ := strconv.Atoi(id)
	svc := services.NewCategoryServices()

	data, err := svc.GetByID(uint(_id))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, utils.SetErrorResponse(http.StatusInternalServerError, "INTERNAL_ERROR", err))
	}

	return ctx.JSON(http.StatusOK, utils.SetSuccessReponse(http.StatusOK, "SUCCESS", data))
}

// GetAll implements contracts.CategoryController.
func (c categoryController) GetAll(ctx echo.Context) error {
	svc := services.NewCategoryServices()
	data, err := svc.GetAll()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, utils.SetErrorResponse(http.StatusInternalServerError, "INTERNAL_ERROR", err))
	}

	return ctx.JSON(http.StatusOK, utils.SetSuccessReponse(http.StatusOK, "SUCCESS", data))
}

// Insert implements contracts.CategoryController.
func (c categoryController) Insert(ctx echo.Context) error {
	var request types.RequestCreateCategory
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse(http.StatusBadRequest, "ERROR_VALIDATION", err))
	}

	svc := services.NewCategoryServices()
	data, err := svc.Insert(request)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, utils.SetErrorResponse(http.StatusInternalServerError, "INTERNAL_ERROR", err))
	}

	return ctx.JSON(http.StatusCreated, utils.SetSuccessReponse(http.StatusCreated, "SUCCESS", data))
}

func NewCategoryController() contracts.CategoryController {
	return categoryController{}
}
