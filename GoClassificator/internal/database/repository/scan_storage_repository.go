package repository

import (
	"GoClassificator/internal/database/models"
)

func CreateScanHistory(databaseID uint, status string) (uint, error) {
	db := GetDB()

	scanHistory := models.ScanHistory{
		DatabaseID: databaseID,
		ScanStatus: status,
	}

	result := db.Create(&scanHistory)
	if result.Error != nil {
		return 0, result.Error
	}

	return scanHistory.ID, nil
}

func UpdateScanHistory(scanID uint, status string) error {
	db := GetDB()

	result := db.Model(&models.ScanHistory{}).Where("id = ?", scanID).Update("scan_status", status)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func SaveScannedTable(scanID uint, tableName string) (uint, error) {
	db := GetDB()

	scannedTable := models.ScannedTable{
		ScanID:    scanID,
		TableName: tableName,
	}

	result := db.Create(&scannedTable)
	if result.Error != nil {
		return 0, result.Error
	}

	return scannedTable.ID, nil
}

// Save the scan result
func SaveScanResult(tableID uint, column string, infoType string) error {
	db := GetDB()

	scanResult := models.ScanResult{
		TableID:         tableID,
		ColumnName:      column,
		InformationType: infoType,
	}

	result := db.Create(&scanResult)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
