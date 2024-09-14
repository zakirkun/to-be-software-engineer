package router

import (
	routeV1 "imzakir.dev/e-commerce/router/v1"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func InitRouters() http.Handler {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"messages": "Hello World!", "request-id": c.Request().Header.Get(echo.HeaderXRequestID)})
	})

	// Versioning
	v1 := e.Group("/v1")
	{
		routeV1.CategoryRoute(v1)
		routeV1.ProductRoute(v1)
		routeV1.CostumerRoute(v1)
	}

	return e
}
