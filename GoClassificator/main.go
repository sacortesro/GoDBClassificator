package main

import (
	"GoClassificator/internal/api"
	"GoClassificator/internal/config"
	"GoClassificator/internal/database/repository"
	"GoClassificator/internal/logger"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	logger.InitLogger("logs/") // Initialize logger

	config.LoadEnv() // Load environment variables

	repository.InitDatabase() // Initialize database

	e := echo.New() // Create an instance of Echo

	// Middleware for logging and error recovery
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	api.SetupRoutes(e) // Setup API routes

	// Start the server on port 8080
	log.Println("Server running at http://localhost:8080")
	e.Logger.Fatal(e.Start(":8080"))
}
