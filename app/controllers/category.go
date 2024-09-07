package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	"imzakir.dev/e-commerce/app/domains/contracts"
	"imzakir.dev/e-commerce/app/domains/types"
	"imzakir.dev/e-commerce/app/services"
	"imzakir.dev/e-commerce/utils"
)

type categoryController struct{}

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
