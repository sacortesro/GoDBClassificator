package repository

import (
	"GoClassificator/internal/database/models"
	"fmt"
)

/*****************************************************
 * 													 *
 * Repository that handles the database scan results *
 * 													 *
 *****************************************************/

// GetLatestCompletedScanID retrieves the latest completed scan ID for a given database ID
func GetLatestCompletedScanID(databaseID uint) (uint, error) {
	db := GetDB()
	var latestScan uint
	if err := db.Model(&models.ScanHistory{}).Select("MAX(id) AS ID").Where("scan_status = ? AND database_id = ?", "COMPLETED", databaseID).Scan(&latestScan).Error; err != nil {
		return 0, fmt.Errorf("error getting latest scan timestamp: %v", err)
	}
	return latestScan, nil
}

// GetScannedTables retrieves the scanned tables for a given scan ID
func GetScannedTables(scanID uint) ([]models.ScannedTable, error) {
	db := GetDB()
	var scannedTables []models.ScannedTable
	if err := db.Where("scan_id = ?", scanID).Find(&scannedTables).Error; err != nil {
		return nil, fmt.Errorf("error getting scanned tables: %v", err)
	}
	return scannedTables, nil
}

// GetScanResults retrieves the scan results for a given list of table IDs
func GetScanResults(tableIDs []uint) ([]models.ScanResult, error) {
	db := GetDB()
	var scanResults []models.ScanResult
	if err := db.Where("table_id IN (?)", tableIDs).Find(&scanResults).Error; err != nil {
		return nil, fmt.Errorf("error getting scan results: %v", err)
	}
	return scanResults, nil
}
