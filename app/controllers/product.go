package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"teukufuad/e-commerce/app/domains/contracts"
	"teukufuad/e-commerce/app/domains/types"
	"teukufuad/e-commerce/app/services"
	"teukufuad/e-commerce/utils"
)

type productController struct{}

func (c productController) Update(ctx echo.Context) error {
	var request types.RequestProduct
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse(http.StatusBadRequest, "ERROR_VALIDATION", err))
	}

	svc := services.NewProductService()
	id, _ := strconv.Atoi(ctx.Param("id"))
	data, err := svc.Update(request, id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, utils.SetErrorResponse(http.StatusInternalServerError, "INTERNAL_ERROR", err))
	}

	return ctx.JSON(http.StatusCreated, utils.SetSuccessReponse(http.StatusCreated, "SUCCESS", data))
}

func (c productController) Delete(ctx echo.Context) error {
	svc := services.NewProductService()
	id, _ := strconv.Atoi(ctx.Param("id"))
	err := svc.Delete(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, utils.SetErrorResponse(http.StatusInternalServerError, "INTERNAL_ERROR", err))
	}

	return ctx.JSON(http.StatusCreated, utils.SetSuccessReponse(http.StatusCreated, "SUCCESS", nil))
}

func (c productController) Get(ctx echo.Context) error {
	svc := services.NewProductService()
	id, _ := strconv.Atoi(ctx.Param("id"))
	data, err := svc.Get(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, utils.SetErrorResponse(http.StatusInternalServerError, "INTERNAL_ERROR", err))
	}

	return ctx.JSON(http.StatusCreated, utils.SetSuccessReponse(http.StatusCreated, "SUCCESS", data))
}

// GetAll implements contracts.CategoryController.
func (c productController) GetAll(ctx echo.Context) error {
	svc := services.NewProductService()
	data, err := svc.GetAll()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, utils.SetErrorResponse(http.StatusInternalServerError, "INTERNAL_ERROR", err))
	}

	return ctx.JSON(http.StatusCreated, utils.SetSuccessReponse(http.StatusCreated, "SUCCESS", data))
}

// Insert implements contracts.CategoryController.
func (c productController) Insert(ctx echo.Context) error {
	var request types.RequestProduct
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse(http.StatusBadRequest, "ERROR_VALIDATION", err))
	}

	svc := services.NewProductService()
	data, err := svc.Insert(request)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, utils.SetErrorResponse(http.StatusInternalServerError, "INTERNAL_ERROR", err))
	}

	return ctx.JSON(http.StatusCreated, utils.SetSuccessReponse(http.StatusCreated, "SUCCESS", data))
}

func NewProductController() contracts.ProductController {
	return productController{}
}
