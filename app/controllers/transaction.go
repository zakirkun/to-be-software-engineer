package controllers

import (
	gojwt "github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
	"log"
	"teukufuad/e-commerce/app/domains/contracts"
)

type TransactionController struct {
}

func (t TransactionController) Order(ctx echo.Context) error {
	userToken := ctx.Get("user").(*gojwt.Token)
	claims := userToken.Claims.(gojwt.MapClaims)
	userID := claims["id"].(string)
	log.Println(userID)
	return nil
}

func NewTransactionController() contracts.TransactionController {
	return TransactionController{}
}
