package repository

import (
	"GoClassificator/internal/database/models"
	"GoClassificator/internal/logger"
	"GoClassificator/internal/security"
	"errors"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// GetDB establishes a connection to the MySQL database and returns the GORM DB instance
func GetScanDB(host string, port int, username, password string) (*gorm.DB, error) {

	logger.Infof("%s:%s@tcp(%s:%d)/?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Errorf("Error connecting to database: %v", err)
		return nil, err
	}
	logger.Info("Successfully connected to scan database")
	return db, nil
}

func GetAllDatabases(db *gorm.DB) ([]string, error) {

	var databases []string
	systemDatabases := []string{"information_schema", "mysql", "performance_schema", "sys"}

	// Execute the query to get all databases excluding system databases
	err := db.Raw("SELECT schema_name FROM information_schema.schemata WHERE schema_name NOT IN ?", systemDatabases).Scan(&databases).Error
	if err != nil {
		return nil, err
	}

	return databases, nil
}

// Check if a database connection exists and save it if it doesn't
func CheckAndSaveDatabaseConnection(host string, port int, username string, password string, dbName string) (uint, error) {
	db := GetDB()
	var connection models.DatabaseConnection

	// Check if the connection already exists
	err := db.Where("host = ? AND port = ? AND dbusername = ? AND dbname = ?", host, port, username, dbName).First(&connection).Error
	if err == nil {
		logger.Info("Connection already exists")
		return connection.ID, nil
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		logger.Errorf("Error checking for connection: %v", err)
		return 0, err
	}

	// Connection does not exist, save it
	return SaveDatabaseConnection(host, port, username, password, dbName)
}

// Save the connection in the database
func SaveDatabaseConnection(host string, port int, username string, password string, dbName string) (uint, error) {
	db := GetDB()

	encryptedPassword, err := security.Encrypt(password)
	if err != nil {
		logger.Errorf("Error encrypting password: %v", err)
		return 0, fmt.Errorf("error encrypting password: %v", err)
	}

	connection := models.DatabaseConnection{
		Host:     host,
		Port:     port,
		Username: username,
		Password: encryptedPassword,
		DbName:   dbName,
	}

	result := db.Create(&connection)
	if result.Error != nil {
		return 0, result.Error
	}

	return connection.ID, nil
}

// Get connection data by ID
func GetDatabaseConnection(id uint) (*models.DatabaseConnection, error) {
	db := GetDB()
	var connection models.DatabaseConnection

	if err := db.First(&connection, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, logger.Errorf("connection with ID %d not found", id)
		}
		return nil, err
	}

	decryptedPassword, err := security.Decrypt(connection.Password)
	if err != nil {
		return nil, fmt.Errorf("error decrypting password: %v", err)
	}
	connection.Password = decryptedPassword

	return &connection, nil
}
