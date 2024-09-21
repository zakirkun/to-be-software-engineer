package v1

import (
	"github.com/labstack/echo"
	"teukufuad/e-commerce/app/controllers"
)

func TransactionRoute(e *echo.Group) {

	// controller
	transactionController := controllers.NewTransactionController()
	//e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
	//	SigningKey:  []byte(config.GetString("jwt.signature_key")),
	//	TokenLookup: "header:Authorization",
	//	AuthScheme:  "Bearer",
	//}))

	transaction := e.Group("/transaction")
	{
		transaction.POST("", transactionController.Order)
	}
}
