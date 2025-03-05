package test

import (
	"GoClassificator/internal/database/repository"
	"GoClassificator/internal/logger"
	"GoClassificator/internal/services"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGetScanResult tests the GetScanResult function
func TestGetScanResult(t *testing.T) {

	repository.InitDatabase()
	logger.InitLogger("logs/")

	// Run the GetScanResult function
	jsonResult, err := services.GetScanResult(3)
	assert.NoError(t, err)

	fmt.Println(jsonResult)
}
