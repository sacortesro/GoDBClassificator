package webrepository

import (
	"GoClassificator/internal/database/models"
	"GoClassificator/internal/database/repository"
	"time"
)

/****************************************************
* 													*
* Repository that handles the rendering of the scan *
* report page 			                            *
* 													*
****************************************************/

// GetScanReportData retrieves the scan report data for a given database connection ID
func GetScanReportData(databaseID uint) (models.ScanReportData, error) {

	db := repository.GetDB()

	var reportData models.ScanReportData

	// Get database connection details
	var dbConnection models.DatabaseConnection
	if err := db.First(&dbConnection, databaseID).Error; err != nil {
		return reportData, err
	}

	// Set basic report data
	reportData.Date = time.Now().Format("January 2, 2006 15:04:05")
	reportData.DatabaseName = dbConnection.DbName
	reportData.Host = dbConnection.Host

	// Get scan count
	var scanCount int64
	if err := db.Model(&models.ScanHistory{}).Where("database_id = ?", databaseID).Count(&scanCount).Error; err != nil {
		return reportData, err
	}
	reportData.ScanCount = int(scanCount)

	// Get total tables and columns
	var totalTables int64
	var totalColumns int64
	if err := db.Model(&models.ScannedTable{}).Where("scan_id IN (SELECT MAX(id) FROM scan_histories WHERE database_id = ? AND scan_status = 'COMPLETED')", databaseID).Count(&totalTables).Error; err != nil {
		return reportData, err
	}
	if err := db.Model(&models.ScanResult{}).Where("table_id IN (SELECT id FROM scanned_tables WHERE scan_id IN (SELECT MAX(id) FROM scan_histories WHERE database_id = ? AND scan_status = 'COMPLETED'))", databaseID).Count(&totalColumns).Error; err != nil {
		return reportData, err
	}
	reportData.TotalTables = int(totalTables)
	reportData.TotalColumns = int(totalColumns)

	// Get data types summary
	var dataTypesSummary []models.DataTypeSummary
	if err := db.Model(&models.ScanResult{}).
		Select("information_type AS type, COUNT(*) AS count").
		Where("table_id IN (SELECT id FROM scanned_tables WHERE scan_id IN (SELECT MAX(id) FROM scan_histories WHERE database_id = ? AND scan_status = 'COMPLETED'))", databaseID).
		Group("information_type").
		Scan(&dataTypesSummary).Error; err != nil {
		return reportData, err
	}
	reportData.DataTypesSummary = dataTypesSummary

	// Get table information
	var scannedTables []models.ScannedTable
	if err := db.Where("scan_id IN (SELECT MAX(id) FROM scan_histories WHERE database_id = ? AND scan_status = 'COMPLETED')", databaseID).Find(&scannedTables).Error; err != nil {
		return reportData, err
	}

	for _, table := range scannedTables {
		var tableInfo models.TableInfo
		tableInfo.Name = table.TableName

		var totalColumns int64
		if err := db.Model(&models.ScanResult{}).Where("table_id = ?", table.ID).Count(&totalColumns).Error; err != nil {
			return reportData, err
		}

		tableInfo.ColumnCount = int(totalColumns)

		reportData.Tables = append(reportData.Tables, tableInfo)
	}

	return reportData, nil
}
