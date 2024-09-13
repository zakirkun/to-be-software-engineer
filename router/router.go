package router

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"imzakir.dev/e-commerce/app/controllers"
	appMiddleware "imzakir.dev/e-commerce/app/middleware"
)

func InitRouters() http.Handler {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	// Auth Handler

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"messages": "Hello World!", "request-id": c.Request().Header.Get(echo.HeaderXRequestID)})
	})

	e.GET("/restricted", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, map[string]string{"messages": "Privated Area!", "request-id": ctx.Request().Header.Get(echo.HeaderXRequestID)})

	}, appMiddleware.AuthMiddleware)

	// controller
	categoryController := controllers.NewCategoryController()
	customerController := controllers.NewCustomerController()
	productController := controllers.NewProductController()

	// Versioning
	v1 := e.Group("/v1")
	{
		category := v1.Group("/category")
		{
			category.POST("/save", categoryController.Insert)
			category.GET("/", categoryController.GetAll)
			category.GET("/:id", categoryController.Show)
			category.PUT("/:id", categoryController.Edit)
			category.DELETE("/:id", categoryController.Delete)
		}
		customer := v1.Group("/auth")
		{
			customer.POST("/register", customerController.Register)
			customer.POST("/login", customerController.Login)
		}
		product := v1.Group("/product")
		{
			product.POST("/save", productController.Create)
			product.GET("/", productController.GetAll)
			product.GET("/:id", productController.GetDetail)
		}
	}

	return e
}
