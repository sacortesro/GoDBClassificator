package auth

import (
	"GoClassificator/internal/database/repository"
	"GoClassificator/internal/logger"
	"errors"
)

/****************************************************
 * API Key model for the database					*
 * 													*
 ****************************************************/

type APIKey struct {
	ID  uint   `gorm:"primaryKey"`
	Key string `gorm:"column:api_key;not null;unique"`
}

// ValidateAPIKey checks if the given API Key exists in the database
func ValidateAPIKey(apiKey string) bool {
	db := repository.GetDB()
	var key APIKey
	err := db.Where("api_key = ?", apiKey).First(&key).Error
	return err == nil // Returns true if API key is found
}

// GenerateAPIKey creates a new API Key and stores it in the database
func GenerateAPIKey(newKey string) error {
	db := repository.GetDB()

	var count int64
	db.Model(&APIKey{}).Where("api_key = ?", newKey).Count(&count)
	if count > 0 {
		logger.Warn("API Key already exists")
		return errors.New("API Key already exists")
	}

	return db.Create(&APIKey{Key: newKey}).Error
}
