package router

import (
	"net/http"

	"github.com/labstack/echo"
)

func InitRouters() http.Handler {

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"messages": "Hello World!", "request-id": c.Request().Header.Get(echo.HeaderXRequestID)})
	})

	return e
}
