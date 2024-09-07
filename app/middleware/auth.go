package middleware

import "github.com/labstack/echo"

func AuthMiddleware(h echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {

		// Todo Implement Auth Middleware

		return nil
	}
}
