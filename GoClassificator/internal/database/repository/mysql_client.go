package repository

import (
	"GoClassificator/internal/config"
	"GoClassificator/internal/logger"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

// DB is the global variable for the MySQL connection
var DB *gorm.DB

// InitDatabase initializes the connection to MySQL
func InitDatabase() {
	// Get credentials from environment variables
	user := config.GetEnv("DB_USER")
	password := config.GetEnv("DB_PASSWORD")
	host := config.GetEnv("DB_HOST")
	port := config.GetEnv("DB_PORT")
	dbName := config.GetEnv("DB_NAME")

	// Build the DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbName)

	// Assign the connection to the global variable
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Info),
	})

	if err != nil {
		logger.Fatalf("Error connectig to database: %v", err)
	}

	// Assign the connection to the global variable
	DB = db
	logger.Info("Successfully connected to database")
}

// GetDB returns the database instance
func GetDB() *gorm.DB {
	return DB
}
