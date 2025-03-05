package test

import (
	"GoClassificator/internal/database/repository"
	"GoClassificator/internal/logger"
	"GoClassificator/internal/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestScanDatabase tests the ScanDatabase function
func TestScanDatabase(t *testing.T) {

	repository.InitDatabase()
	logger.InitLogger("logs/")

	// Run the ScanDatabase function
	err := services.ScanDatabase(3)
	assert.NoError(t, err)
}
