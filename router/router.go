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
	productController := controllers.NewProductController()
	customerController := controllers.NewCustomerController()
	transactionController := controllers.NewTransactionController()
	// Versioning

	v1 := e.Group("/v1")
	{
		category := v1.Group("/category")
		{
			category.POST("/save", categoryController.Insert)
			category.GET("/", categoryController.GetAll)
			category.GET("/:id", categoryController.Show)
			category.DELETE("/:id", categoryController.Delete)
			category.PUT("/:id", categoryController.Edit)
			category.GET("/detail/:id", categoryController.GetByID)
		}

		product := v1.Group("/product")
		{ 
			product.POST("/save", productController.Insert)
			product.GET("/", productController.GetAll)
			product.GET("/:id", productController.Show)
			product.DELETE("/:id", productController.Delete)
			product.PUT("/:id", productController.Edit)
		}

		customer := v1.Group("/customer")
		{
			customer.POST("/save", customerController.Insert)
			customer.GET("/:id",customerController.Show)
		}

		transaction := v1.Group("/transaction")
		{
			transaction.POST("/order",transactionController.Insert)
			transaction.GET("/:id", transactionController.Show)
		}
		
	}

	return e
}
