package middleware

import (
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"imzakir.dev/e-commerce/pkg/config"
	"imzakir.dev/e-commerce/pkg/jwt"
)

func AuthMiddleware(h echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {

		log.Info("Hit Middleware")

		// Todo Implement Auth Middleware
		// Get the Authorization header
		authHeader := ctx.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Missing or invalid token")
		}

		// Split the header to extract the token (format: "Bearer <token>")
		splitToken := strings.Split(authHeader, "Bearer ")
		if len(splitToken) != 2 {
			return echo.NewHTTPError(http.StatusUnauthorized, "Malformed token")
		}

		tokenString := splitToken[1]
		_jwt := jwt.NewJWTImpl(config.GetString("jwt.signature_key"), config.GetInt("jwt.day_expired"))
		valid, err := _jwt.ValidateToken(tokenString)
		if err != nil || !valid {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}

		parse, _ := _jwt.ParseToken(tokenString)
		ctx.Set("username", parse["username"])

		return h(ctx)
	}
}
