package middleware

import (
	"net/http"
	"strings"

	"github.com/LanternNassi/IMSController/internal/utils"

	"github.com/labstack/echo"
)

func AuthenticationMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tokenString := c.Request().Header.Get("Authorization")
			if tokenString == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Missing authentication token"})
			}

			tokenParts := strings.Split(tokenString, " ")
			if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid authentication token"})
			}

			tokenString = tokenParts[1]

			claims, err := utils.VerifyToken(tokenString)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid authentication token"})
			}

			c.Set("user_id", claims["user_id"])
			return next(c)
		}
	}
}
