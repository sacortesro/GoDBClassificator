package services

import (
	"GoClassificator/internal/database/models"
	"GoClassificator/internal/database/repository"
	"GoClassificator/internal/logger"
)

/****************************************************
 * 													*
 * Service that returns the scan result information *
 * 													*
 ****************************************************/

// Structs for JSON response
type Column struct {
	ColumnName      string `json:"column_name"`
	InformationType string `json:"information_type"`
}

type Table struct {
	TableName string   `json:"table_name"`
	Columns   []Column `json:"columns"`
}

type Schema struct {
	SchemaName string  `json:"schema_name"`
	Tables     []Table `json:"tables"`
}

type DatabaseStructure struct {
	DatabaseName string   `json:"database_name"`
	Schemas      []Schema `json:"schemas"`
}

// Get scan results by database ID
func GetScanResult(id uint) (DatabaseStructure, error) {

	logger.Info("Getting scan results")

	// Query the database connection
	connection, err := repository.GetDatabaseConnection(id)
	if err != nil {
		return DatabaseStructure{}, err
	}

	// Query the latest completed scan timestamp and scan ID
	latestScan, err := repository.GetLatestCompletedScanID(id)
	if err != nil {
		return DatabaseStructure{}, err
	}

	// Query the scanned tables and scan results for the latest scan
	scannedTables, err := repository.GetScannedTables(latestScan)
	if err != nil {
		return DatabaseStructure{}, err
	}

	scanResults, err := repository.GetScanResults(getTableIDs(scannedTables))
	if err != nil {
		return DatabaseStructure{}, err
	}
	// Create a map to store the database structure
	databaseStructure := DatabaseStructure{
		DatabaseName: connection.DbName,
		Schemas:      []Schema{},
	}

	// store tables by schema
	tablesBySchema := make(map[string][]Table)

	// Iterate over the scanned tables and build the JSON structure
	for _, table := range scannedTables {
		// Create a table struct
		newTable := Table{
			TableName: table.TableName,
			Columns:   []Column{},
		}

		// Add columns to the table
		for _, result := range scanResults {
			if result.TableID == table.ID {
				column := Column{
					ColumnName:      result.ColumnName,
					InformationType: result.InformationType,
				}
				newTable.Columns = append(newTable.Columns, column)
			}
		}

		// Add the table to the schema
		tablesBySchema["default"] = append(tablesBySchema["default"], newTable)
	}

	// Add the tables to the schema
	for schemaName, tables := range tablesBySchema {
		schema := Schema{
			SchemaName: schemaName,
			Tables:     tables,
		}
		databaseStructure.Schemas = append(databaseStructure.Schemas, schema)
	}
	logger.Info("Scan results retrieved")

	return databaseStructure, nil
}

// Helper function to get table IDs from scanned tables
func getTableIDs(tables []models.ScannedTable) []uint {
	ids := make([]uint, len(tables))
	for i, table := range tables {
		ids[i] = table.ID
	}
	return ids
}
