package controllers

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"imzakir.dev/e-commerce/app/domains/contracts"
	"imzakir.dev/e-commerce/app/domains/types"
	"imzakir.dev/e-commerce/app/services"
	"imzakir.dev/e-commerce/utils"
)

type orderController struct{}

// HandleCallback implements contracts.OrderController.
func (o orderController) HandleCallback(ctx echo.Context) error {
	panic("unimplemented")
}

// CreateTransaction implements contracts.OrderController.
func (o orderController) CreateTransaction(ctx echo.Context) error {
	var request types.RequestCreateTransaction
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse(http.StatusBadRequest, "ERROR_VALIDATION", err))
	}

	username, _ := ctx.Get("username").(string)

	log.Println(ctx.Get("username"))

	// assign username
	request.Username = username

	svc := services.NewOrderServices()
	data, err := svc.CreateTransaction(request)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, utils.SetErrorResponse(http.StatusInternalServerError, "INTERNAL_ERROR", err))
	}

	return ctx.JSON(http.StatusCreated, utils.SetSuccessReponse(http.StatusCreated, "SUCCESS", data))
}

// GetTransaction implements contracts.OrderController.
func (o orderController) GetTransaction(ctx echo.Context) error {
	id := ctx.Param("id")
	if id == "" {
		return ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse(http.StatusBadRequest, "ERROR_VALIDATION", errors.New("invalid id")))
	}
	_id, _ := strconv.Atoi(id)

	svc := services.NewOrderServices()
	data, err := svc.GetTransaction(uint(_id))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, utils.SetErrorResponse(http.StatusInternalServerError, "INTERNAL_ERROR", err))
	}

	return ctx.JSON(http.StatusCreated, utils.SetSuccessReponse(http.StatusCreated, "SUCCESS", data))
}

func NewOrderController() contracts.OrderController {
	return orderController{}
}
