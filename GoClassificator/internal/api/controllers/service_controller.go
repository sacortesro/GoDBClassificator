package controllers

import (
	"GoClassificator/internal/logger"
	"GoClassificator/internal/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

/****************************************************
 * 													*
 * Controller that handles the http request related *
 * to the database connection and scans.			*
 * 													*
 ****************************************************/

// Structure for the connection request
type DatabaseRequest struct {
	Host     string `json:"host" validate:"required"`
	Port     int    `json:"port" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// Create a new database connection
func CreateDatabase(c echo.Context) error {

	logger.Info("Creating database connection")

	var req DatabaseRequest

	// Validate input JSON
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid data"})
	}

	// Save the connection in the database
	id, err := services.SaveDatabaseConnection(req.Host, req.Port, req.Username, req.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	logger.Infof("Connection created with ID: %d", id)

	return c.JSON(http.StatusCreated, map[string]any{"id": id})
}

// Execute the database scan
func ScanDatabase(c echo.Context) error {
	id := c.Param("id")

	// Convert id to uint
	uintID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	// Execute the scan
	err = services.ScanDatabase(uint(uintID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "Scan succesfull"})
}

// Get the result of a scan
func GetScanResult(c echo.Context) error {
	id := c.Param("id")

	// Convert id to uint
	uintID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	result, err := services.GetScanResult(uint(uintID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
