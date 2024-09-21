package controllers

import (
	"github.com/labstack/echo"
	"net/http"
	"teukufuad/e-commerce/app/domains/contracts"
	"teukufuad/e-commerce/app/domains/types"
	"teukufuad/e-commerce/app/services"
	"teukufuad/e-commerce/utils"
)

type TransactionController struct {
}

func (t TransactionController) Order(ctx echo.Context) error {
	username, _ := ctx.Get("username").(string)

	var request types.RequestOrder
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse(http.StatusBadRequest, "ERROR_VALIDATION", err))
	}

	request.Username = username
	svc := services.NewTransactionService()
	data, err := svc.Order(request)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, utils.SetErrorResponse(http.StatusInternalServerError, "INTERNAL_ERROR", err))
	}

	return ctx.JSON(http.StatusCreated, utils.SetSuccessReponse(http.StatusCreated, "SUCCESS", data))
}

func NewTransactionController() contracts.TransactionController {
	return TransactionController{}
}
