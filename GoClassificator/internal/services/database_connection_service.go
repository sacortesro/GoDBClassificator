package services

import (
	"GoClassificator/internal/database/repository"
	"GoClassificator/internal/logger"
)

/****************************************************
 * 													*
 * Service that handles the database connection to  *
 * store											*
 * 													*
 ****************************************************/

// SaveDatabaseConnection saves the database connection in the database
func SaveDatabaseConnection(host string, port int, username string, password string) ([]uint, error) {

	// Get the database connection to scan. This ensures that the connection is valid
	logger.Info("Getting scan connection database")
	db, err := repository.GetScanDB(host, port, username, password)
	if err != nil {
		logger.Errorf("Error connecting to database: %v", err)
		return nil, err
	}

	// Get the databases in the server
	databases, err := repository.GetAllDatabases(db)
	if err != nil {
		logger.Errorf("Error retrieving databases: %v", err)
		return nil, err
	}

	var ids []uint

	// Save the connection for each database
	for _, dbName := range databases {
		id, err := repository.CheckAndSaveDatabaseConnection(host, port, username, password, dbName)
		if err != nil {
			logger.Fatalf("Error saving connection: %v", err)
			return nil, err
		}
		ids = append(ids, id)
	}

	return ids, nil
}
