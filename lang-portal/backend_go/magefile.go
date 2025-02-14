//go:build mage

package main

import (
	"fmt"
	"os"
	"path/filepath"

	"backend_go/internal/db"
	"backend_go/internal/models"
)

const (
	dbPath     = "words.db"
	migrations = "db/migrations"
)

// Init initializes the database
func Init() error {
	fmt.Println("Initializing database...")

	// Create db directory if it doesn't exist
	if err := os.MkdirAll(filepath.Dir(dbPath), 0755); err != nil {
		return fmt.Errorf("failed to create db directory: %w", err)
	}

	// Initialize database connection
	if err := models.InitDB(dbPath); err != nil {
		return fmt.Errorf("failed to initialize database: %w", err)
	}

	return nil
}

// Migrate runs database migrations
func Migrate() error {
	fmt.Println("Running migrations...")

	if err := models.InitDB(dbPath); err != nil {
		return fmt.Errorf("failed to initialize database: %w", err)
	}

	manager := db.NewMigrationManager(migrations)
	if err := manager.RunMigrations(); err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	return nil
}

// Reset clears all data and reinitializes the database
func Reset() error {
	fmt.Println("Resetting database...")

	// Remove existing database
	if err := os.Remove(dbPath); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to remove database: %w", err)
	}

	// Initialize new database
	if err := Init(); err != nil {
		return err
	}

	// Run migrations
	if err := Migrate(); err != nil {
		return err
	}

	return nil
}

// ResetHistory clears only study history while preserving vocabulary and groups
func ResetHistory() error {
	fmt.Println("Resetting study history...")

	if err := models.InitDB(dbPath); err != nil {
		return fmt.Errorf("failed to initialize database: %w", err)
	}

	// Delete all study related records
	if err := models.DB.Exec("DELETE FROM word_review_items").Error; err != nil {
		return fmt.Errorf("failed to delete word reviews: %w", err)
	}
	if err := models.DB.Exec("DELETE FROM study_sessions").Error; err != nil {
		return fmt.Errorf("failed to delete study sessions: %w", err)
	}

	return nil
}

// Dev starts the development server
func Dev() error {
	fmt.Println("Starting development server...")
	// TODO: Implement development server startup
	return nil
}

// Seed imports seed data
func Seed() error {
	fmt.Println("Seeding database...")

	if err := models.InitDB(dbPath); err != nil {
		return fmt.Errorf("failed to initialize database: %w", err)
	}

	manager := db.NewSeedManager("db/seeds")

	// Define seed files and their corresponding groups
	seeds := map[string]string{
		"basic_greetings.json": "Basic Greetings",
		"common_nouns.json":    "Common Nouns",
		"basic_verbs.json":     "Basic Verbs",
	}

	for file, group := range seeds {
		fmt.Printf("Loading %s into group %s...\n", file, group)
		if err := manager.LoadSeedFile(file, group); err != nil {
			return fmt.Errorf("failed to load seed file %s: %w", file, err)
		}
	}

	return nil
}
