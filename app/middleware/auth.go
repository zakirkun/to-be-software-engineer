package middleware

import (
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func AuthMiddleware(h echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {

		log.Info("Hit Middleware")

		// Todo Implement Auth Middleware

		return nil
	}
}
