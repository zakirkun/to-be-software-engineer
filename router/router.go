package router

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"imzakir.dev/e-commerce/app/controllers"
)

func InitRouters() http.Handler {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"messages": "Hello World!", "request-id": c.Request().Header.Get(echo.HeaderXRequestID)})
	})

	// controller
	categoryController := controllers.NewCategoryController()

	// Versioning
	v1 := e.Group("/v1")
	{
		category := v1.Group("/category")
		{
			category.POST("", categoryController.Insert)
			category.GET("", categoryController.GetAll)
			category.GET("/:id", categoryController.Get)
		}
	}

	return e
}
