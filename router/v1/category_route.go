package v1

import (
	"github.com/labstack/echo"
	"teukufuad/e-commerce/app/controllers"
)

func CategoryRoute(e *echo.Group) {
	// controller
	categoryController := controllers.NewCategoryController()

	category := e.Group("/category")
	{
		category.POST("", categoryController.Insert)
		category.GET("", categoryController.GetAll)
		category.GET("/:id", categoryController.Get)
		category.PUT("/:id", categoryController.Update)
		category.DELETE("/:id", categoryController.Delete)
	}
}
