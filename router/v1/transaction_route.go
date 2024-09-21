package v1

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
	"teukufuad/e-commerce/app/controllers"
	"teukufuad/e-commerce/pkg/config"
)

func TransactionRoute(e *echo.Group) {

	// controller

	transactionController := controllers.NewTransactionController()
	log.Println("tessssssssssssssssssssssssssssssssss ", config.GetString("jwt.signature_key"))
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  []byte(config.GetString("jwt.signature_key")),
		TokenLookup: "header:Authorization",
		AuthScheme:  "Bearer",
	}))

	transaction := e.Group("/transaction")
	{
		transaction.POST("", transactionController.Order)
	}
}
