package middleware

import (
	"GoClassificator/internal/api/auth"
	"GoClassificator/internal/logger"
	"net/http"

	"github.com/labstack/echo/v4"
)

/****************************************************
 * 													*
 * Middleware that handles the API Key validation	*
 * 													*
 * 													*
 ****************************************************/

// APIKeyMiddleware validates the API Key in the request header
func APIKeyMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		apiKey := c.Request().Header.Get("X-API-Key")

		// Validate API Key
		if !auth.ValidateAPIKey(apiKey) {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid API Key"})
		}

		// Proceed to the next handler
		logger.Info("API Key validated")
		return next(c)
	}
}
