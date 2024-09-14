package v1

import (
	"github.com/labstack/echo"
	"imzakir.dev/e-commerce/app/controllers"
)

func CostumerRoute(e *echo.Group) {
	// controller
	categoryController := controllers.NewCostumerController()

	category := e.Group("/customer")
	{
		category.POST("/signin", categoryController.SignIn)
		category.POST("/signup", categoryController.SignUp)
	}
}
