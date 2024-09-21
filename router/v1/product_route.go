package v1

import (
	"github.com/labstack/echo"
	"teukufuad/e-commerce/app/controllers"
)

func ProductRoute(e *echo.Group) {

	productController := controllers.NewProductController()

	product := e.Group("/product")
	{
		product.POST("", productController.Insert)
		product.GET("", productController.GetAll)
		product.GET("/:id", productController.Get)
		product.PUT("/:id", productController.Update)
		product.DELETE("/:id", productController.Delete)
	}
}
