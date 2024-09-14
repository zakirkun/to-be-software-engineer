package controllers

import (
	"github.com/labstack/echo"
	"imzakir.dev/e-commerce/app/domains/contracts"
	"imzakir.dev/e-commerce/app/domains/types"
	"imzakir.dev/e-commerce/app/services"
	"imzakir.dev/e-commerce/pkg/config"
	"imzakir.dev/e-commerce/pkg/jwt"
	"imzakir.dev/e-commerce/utils"
	"net/http"
	"strconv"
	"strings"
)

type transactionController struct{}

func (t transactionController) Create(ctx echo.Context) error {
	tokenString := ctx.Request().Header.Get("Authorization")
	signature := config.GetString("jwt.signature_key")
	expired, _ := strconv.Atoi(config.GetString("jwt.day_expired"))

	jwtClaims := jwt.NewJWTImpl(signature, expired)
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	_, err := jwtClaims.ValidateToken(tokenString)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, utils.SetErrorResponse(http.StatusUnauthorized, "Unauthorized", err))
	}

	tokenParse, err := jwtClaims.ParseToken(tokenString)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, utils.SetErrorResponse(http.StatusUnauthorized, "Unauthorized", err))
	}
	username := tokenParse["username"].(string)

	var request types.RequestCreateTransaction
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, utils.SetErrorResponse(http.StatusBadRequest, "ERROR_VALIDATION", err))
	}

	svc := services.NewTransactionService()
	data, err := svc.AddTransaction(request, username)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, utils.SetErrorResponse(http.StatusInternalServerError, "INTERNAL_ERROR", err))
	}

	return ctx.JSON(http.StatusCreated, utils.SetSuccessReponse(http.StatusCreated, "SUCCESS ADD PRODUCT", data))
}

func NewTransactionController() contracts.TransactionController {
	return transactionController{}
}
