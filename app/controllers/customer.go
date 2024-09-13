package controllers

import (
	"github.com/labstack/echo"
	"imzakir.dev/e-commerce/app/domains/contracts"
	"imzakir.dev/e-commerce/app/domains/types"
	"imzakir.dev/e-commerce/app/services"
	"imzakir.dev/e-commerce/utils"
	"net/http"
)

type customerController struct{}

var svc = services.NewCustomerServices()

func (c customerController) Register(ctx echo.Context) error {
	var request types.RequestCreateCustomer

	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse(http.StatusBadRequest, "ERROR_VALIDATION", err))
	}

	data, err := svc.AddCustomer(request)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, utils.SetErrorResponse(http.StatusInternalServerError, "INTERNAL_ERROR", err))
	}

	return ctx.JSON(http.StatusCreated, utils.SetSuccessReponse(http.StatusCreated, "BERHASIL REGISTER", data))
}

func NewCustomerController() contracts.CustomerController {
	return customerController{}
}
