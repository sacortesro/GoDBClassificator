package test

import (
	"GoClassificator/internal/database/repository"
	"GoClassificator/internal/logger"
	"GoClassificator/internal/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test to save a MySQL connection
func TestSaveDatabaseConnection(t *testing.T) {

	logger.InitLogger("logs/")
	repository.InitDatabase()

	// Attempt to save a connection
	id, err := services.SaveDatabaseConnection("localhost", 3307, "root", "1234")

	// Verify that there are no errors
	assert.NoError(t, err)
	assert.NotZero(t, id, "The connection ID should not be 0")

}
