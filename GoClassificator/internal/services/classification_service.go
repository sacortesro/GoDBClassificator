package services

import (
	"GoClassificator/internal/database/repository"
	"GoClassificator/internal/logger"
	"fmt"
	"log"
	"regexp"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/****************************************************
 * 													*
 * Service that handles the database classification *
 * process											*
 * 													*
 ****************************************************/

// Mapping of sensitive data types
var sensitiveDataPatterns = map[string]string{
	"USERNAME":   `(?i)username|user_name`,
	"FIRST_NAME": `(?i)first_name|fname`,
	"LAST_NAME":  `(?i)last_name|lname`,
	"GENDER":     `(?i)gender`,
	"EMPLOYEE":   `(?i)emp_no`,
	"DATE":       `(?i)date`,
	"DEPARTMENT": `(?i)dept`,
}

// ScanDatabase Execute scan on a database
func ScanDatabase(id uint) error {

	logger.Info("Scanning database")
	// Get connection data
	connection, err := repository.GetDatabaseConnection(id)
	if err != nil {
		return err
	}

	// Connect to the MySQL database to scan
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		connection.Username, connection.Password, connection.Host, connection.Port, connection.DbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("error connecting to the database: %v", err)
	}

	// Get the tables
	var tables []string
	db.Raw("SELECT table_name FROM information_schema.tables WHERE table_schema = ?", connection.DbName).Scan(&tables)

	// Save the scan history
	scanID, err := repository.CreateScanHistory(id, "IN_PROGRESS")
	if err != nil {
		return fmt.Errorf("error saving scan history: %v", err)
	}

	// Analyze each table and its columns
	for _, table := range tables {
		log.Printf("Analyzing table: %s\n", table)

		// Save the scanned table
		table_id, err := repository.SaveScannedTable(scanID, table)
		if err != nil {
			log.Printf("Error saving scanned table %s: %v\n", table, err)
			repository.UpdateScanHistory(scanID, "ERROR")
		}

		var columns []string
		db.Raw("SELECT column_name FROM information_schema.columns WHERE table_schema = ? AND table_name = ?", connection.DbName, table).Scan(&columns)

		for _, column := range columns {
			classification := classifyColumn(column)
			log.Printf("   - Column: %s → %s\n", column, classification)

			// Save the scan result
			err := repository.SaveScanResult(table_id, column, classification)
			if err != nil {
				log.Printf("Error saving scan result for column %s: %v\n", column, err)
				repository.UpdateScanHistory(scanID, "ERROR")
			}
		}
	}
	repository.UpdateScanHistory(scanID, "COMPLETED")

	logger.Info("Scan completed")
	return nil
}

// Classify a column based on its name
func classifyColumn(columnName string) string {
	for dataType, pattern := range sensitiveDataPatterns {
		matched, _ := regexp.MatchString(pattern, columnName)
		if matched {
			return dataType
		}
	}
	return "N/A"
}
