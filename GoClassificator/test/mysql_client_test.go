package test

import (
	"GoClassificator/internal/database/repository"
	"reflect"
	"testing"

	"gorm.io/gorm"
)

// TestInitDatabase tests the InitDatabase function
func TestInitDatabase(t *testing.T) {

	// Initialize the database
	repository.InitDatabase()

	// Get the database connection
	DB := repository.GetDB()

	// Check if the DB variable is not nil
	if DB == nil {
		t.Fatal("DB is nil, expected a valid database connection")
	}

	// Check if the DB variable is of type *gorm.DB
	if reflect.TypeOf(DB) != reflect.TypeOf(&gorm.DB{}) {
		t.Fatalf("DB is not of type *gorm.DB, got %T", DB)
	}
}

// TestGetDB tests the GetDB function
func TestGetDB(t *testing.T) {
	// Initialize the database
	repository.InitDatabase()

	// Get the database instance
	db := repository.GetDB()

	// Check if the returned database instance is not nil
	if db == nil {
		t.Fatal("GetDB returned nil, expected a valid database instance")
	}

	// Check if the returned database instance is of type *gorm.DB
	if (reflect.TypeOf(db) != reflect.TypeOf(&gorm.DB{})) {
		t.Fatalf("GetDB returned a value that is not of type *gorm.DB, got %T", db)
	}
}
