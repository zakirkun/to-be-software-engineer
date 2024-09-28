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
	e.Use(middleware.RequestID())
	// Enable Logging Middleware
	e.Use(appMiddleware.Logging)

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
	orderController := controllers.NewOrderController()

	// Versioning
	v1 := e.Group("/v1")
	{

		order := v1.Group("/order", appMiddleware.AuthMiddleware)
		{
			order.GET("/dummy", func(ctx echo.Context) error {
				return ctx.JSON(http.StatusOK, echo.Map{
					"message": "order up",
				})
			})

			order.POST("/", orderController.CreateTransaction)
			order.GET("/status/:id", orderController.GetTransaction)
		}

		auth := v1.Group("/auth")
		{
			auth.POST("/login", customerController.Login)
			auth.POST("/register", customerController.Register)
		}

		category := v1.Group("/category")
		{
			category.POST("/save", categoryController.Insert)
			category.GET("/", categoryController.GetAll)
			category.GET("/:id", categoryController.GetByID)
			category.PUT("/:id", categoryController.Update)
			category.DELETE("/:id", categoryController.Delete)
			category.GET("/search", categoryController.Pagination)
		}

		product := v1.Group("/product")
		{
			product.GET("/category/:id", categoryController.GetByCategory)
			product.POST("/save", productController.Insert)
			product.GET("/", productController.GetAll)
			product.GET("/:id", productController.GetByID)
			product.PUT("/:id", productController.Update)
			product.DELETE("/:id", productController.Delete)
			product.GET("/search", productController.Pagination)
		}
	}

	return e
}
