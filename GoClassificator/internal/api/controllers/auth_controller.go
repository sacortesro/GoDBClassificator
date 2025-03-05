package controllers

import (
	"GoClassificator/internal/api/auth"
	"net/http"

	"github.com/labstack/echo/v4"
)

/****************************************************
 * 													*
 * Controller that handles the authentication		*
 * 													*
 ****************************************************/

// APIKeyRequest represents an API key creation request
type APIKeyRequest struct {
	Key string `json:"key" validate:"required"`
}

// CreateAPIKey generates a new API Key
func CreateAPIKey(c echo.Context) error {
	var req APIKeyRequest

	// Parse JSON input
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	// Generate and store API Key
	err := auth.GenerateAPIKey(req.Key)
	if err != nil {
		return c.JSON(http.StatusConflict, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "API Key created"})
}
