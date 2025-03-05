package api

import (
	"GoClassificator/internal/api/controllers"
	"GoClassificator/internal/api/middleware"

	"github.com/labstack/echo/v4"
)

/****************************************************
 * 													*
 * API routes configuration							*
 * 													*
 * 													*
 ****************************************************/

// SetupRoutes configures the API endpoints
func SetupRoutes(e *echo.Echo) {
	v1 := e.Group("/api/v1/database")

	v1.Use(middleware.APIKeyMiddleware)

	v1.POST("/", controllers.CreateDatabase)       // Register MySQL connection
	v1.POST("/scan/:id", controllers.ScanDatabase) // Execute scan
	v1.GET("/scan/:id", controllers.GetScanResult) // Get scan result

	e.POST("/apikey", controllers.CreateAPIKey)             // Create API Key
	e.GET("/view/report/:id", controllers.RenderScanReport) // Render HTML page
}
