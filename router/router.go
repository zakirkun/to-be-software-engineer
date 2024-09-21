package router

import (
	"net/http"
	routeV1 "teukufuad/e-commerce/router/v1"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func InitRouters() http.Handler {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.Use(middleware.RequestID())

	// Auth Handler

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"messages": "Hello World!", "request-id": c.Request().Header.Get(echo.HeaderXRequestID)})
	})

	// Versioning
	v1 := e.Group("/v1")
	{
		routeV1.CategoryRoute(v1)
		routeV1.ProductRoute(v1)
		routeV1.CostumerRoute(v1)
		routeV1.TransactionRoute(v1)
	}

	return e
}
