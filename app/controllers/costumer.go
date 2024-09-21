package controllers

import (
	"github.com/labstack/echo"
	"net/http"
	"teukufuad/e-commerce/app/domains/contracts"
	"teukufuad/e-commerce/app/domains/types"
	"teukufuad/e-commerce/app/services"
	"teukufuad/e-commerce/utils"
)

type CostumerController struct {
}

func (c CostumerController) SignIn(ctx echo.Context) error {
	var request types.RequestSignIn
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse(http.StatusBadRequest, "ERROR_VALIDATION", err))
	}

	svc := services.NewCustomerService()
	data, err := svc.SignIn(request)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, utils.SetErrorResponse(http.StatusInternalServerError, "INTERNAL_ERROR", err))
	}

	return ctx.JSON(http.StatusCreated, utils.SetSuccessReponse(http.StatusCreated, "SUCCESS", data))
}

func (c CostumerController) SignUp(ctx echo.Context) error {
	var request types.RequestSignUp
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse(http.StatusBadRequest, "ERROR_VALIDATION", err))
	}

	svc := services.NewCustomerService()
	data, err := svc.SignUp(request)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, utils.SetErrorResponse(http.StatusInternalServerError, "INTERNAL_ERROR", err))
	}

	return ctx.JSON(http.StatusCreated, utils.SetSuccessReponse(http.StatusCreated, "SUCCESS", data))
}

func NewCostumerController() contracts.CustomerController {
	return CostumerController{}
}
