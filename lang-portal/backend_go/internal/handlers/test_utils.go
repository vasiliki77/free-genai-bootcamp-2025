package handlers

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/lang-portal/backend_go/internal/models"
	_ "github.com/mattn/go-sqlite3" // Import sqlite3 driver
)

// Helper function to run migrations for tests
func runMigrationsForTest(t *testing.T, db *models.DB) {
	log.Println("Starting migrations in runMigrationsForTest")
	migration, err := os.ReadFile(filepath.Join("..", "..", "db", "migrations", "0001_init.sql")) // Adjust path to migration file
	if err != nil {
		log.Printf("Error reading migration file: %v", err)
		t.Fatalf("Failed to read migration file: %v", err)
	}
	log.Printf("Migration SQL: %s", string(migration)) // Log migration SQL

	_, err = db.Exec(string(migration))
	if err != nil {
		log.Printf("Error executing migration: %v", err)
		log.Printf("Underlying error: %+v", err)
		fmt.Printf("Migration SQL causing error:\n%s\n", string(migration))
		t.Fatalf("Failed to execute migration: %v", err)
	}
	log.Println("Migrations executed successfully in runMigrationsForTest")
}

// SetupTest initializes the test database and runs migrations.
func SetupTest(t *testing.T) *models.DB {
	db, err := models.NewDB(":memory:")
	if err != nil {
		t.Fatalf("Failed to initialize test database: %v", err)
	}
	runMigrationsForTest(t, db)
	return db
}

// TeardownTest closes the test database.
func TeardownTest(db *models.DB) {
	if db != nil {
		db.Close()
	}
}
