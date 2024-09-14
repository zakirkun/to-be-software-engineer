package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"imzakir.dev/e-commerce/app/domains/contracts"
	"imzakir.dev/e-commerce/app/domains/types"
	"imzakir.dev/e-commerce/app/services"
	"imzakir.dev/e-commerce/utils"
)

type productController struct{}

// GetAll implements contracts.CategoryController.
func (c productController) GetAll(ctx echo.Context) error {
	svc := services.NewProductServices()
	data, err := svc.GetAll()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, utils.SetErrorResponse(http.StatusInternalServerError, "INTERNAL_ERROR", err))
	}

	return ctx.JSON(http.StatusOK, utils.SetSuccessReponse(http.StatusOK, "SUCCESS", data))
}

// Insert implements contracts.CategoryController.
func (c productController) Insert(ctx echo.Context) error {
	var request types.RequestCreateProduct
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse(http.StatusBadRequest, "ERROR_VALIDATION", err))
	}

	svc := services.NewProductServices()
	data, err := svc.Insert(request)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, utils.SetErrorResponse(http.StatusInternalServerError, "INTERNAL_ERROR", err))
	}

	return ctx.JSON(http.StatusCreated, utils.SetSuccessReponse(http.StatusCreated, "SUCCESS", data))
}

//get category by id
func (c productController) Show(ctx echo.Context) error {
	paramId := ctx.Param("id") 
	id, err := strconv.Atoi(paramId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, utils.SetErrorResponse(http.StatusInternalServerError, "INTERNAL_ERROR", err))
	}

	data, err := services.NewProductServices().GetProductById(id)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, utils.SetErrorResponse(http.StatusInternalServerError, "INTERNAL_ERROR", err))
	}

	return ctx.JSON(http.StatusOK, utils.SetSuccessReponse(http.StatusOK, "SUCCESS", data))

}

//delete by id
func (c productController) Delete(ctx echo.Context) error  {
	paramId := ctx.Param("id")
	id, err := strconv.Atoi(paramId)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, utils.SetErrorResponse(http.StatusInternalServerError, "INTERNAL_ERROR", err))
	}

	
	data, err := services.NewProductServices().GetProductById(id)

	log.Println(data)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, utils.SetErrorResponse(http.StatusInternalServerError,"INTERNAL_ERROR", err))
	}

	  services.NewProductServices().DeleteId(id)

	return ctx.JSON(http.StatusOK, utils.SetSuccessDeleteReponse(http.StatusOK, "SUCCESS"))
	
	
}

func NewProductController() contracts.ProductController {
	return productController{}
}


