package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"imzakir.dev/e-commerce/app/domains/contracts"
	"imzakir.dev/e-commerce/app/domains/types"
	"imzakir.dev/e-commerce/app/services"
	"imzakir.dev/e-commerce/utils"
)

type categoryController struct{}

// Delete implements contracts.CategoryController.
func (c categoryController) Delete(ctx echo.Context) error {
	paramId := ctx.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse(http.StatusBadRequest, "BAD REQUEST", err))
	}

	svc := services.NewCategoryServices()

	data, err := svc.Delete(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, utils.SetErrorResponse(http.StatusInternalServerError, "INTERNAL_ERROR", err))
	}

	return ctx.JSON(http.StatusOK, utils.SetSuccessReponse(http.StatusOK, "SUCCESS DELETE CATEGORY", data))
}

// Update implements contracts.CategoryController.
func (c categoryController) Update(ctx echo.Context) error {
	paramId := ctx.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse(http.StatusBadRequest, "BAD REQUEST", err))
	}

	var request types.RequestCreateCategory
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse(http.StatusBadRequest, "ERROR_VALIDATION", err))
	}

	svc := services.NewCategoryServices()

	data, err := svc.Update(id, request)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, utils.SetErrorResponse(http.StatusInternalServerError, "INTERNAL_ERROR", err))
	}

	return ctx.JSON(http.StatusOK, utils.SetSuccessReponse(http.StatusOK, "SUCCESS UPDATE CATEGORY", data))
}

// GetDetail implements contracts.CategoryController.
func (c categoryController) GetDetail(ctx echo.Context) error {
	svc := services.NewCategoryServices()
	id, _ := strconv.Atoi(ctx.Param("id"))
	data, err := svc.GetDetail(id)

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
