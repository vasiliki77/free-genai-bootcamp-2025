//go:build mage
package main

import (
	"database/sql"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	_ "github.com/mattn/go-sqlite3"
)

// Dev starts the development server
func Dev() error {
	env := os.Environ()
	env = append(env, "GO111MODULE=on")
	fmt.Println("Starting server...")
	cmd := exec.Command("go", "run", "cmd/server/main.go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = env
	return cmd.Run()
	  }

// InitDB initializes the SQLite database
func InitDB() error {
	dbPath := "words.db"
	
	// Remove existing database if it exists
	if _, err := os.Stat(dbPath); err == nil {
		if err := os.Remove(dbPath); err != nil {
			return fmt.Errorf("failed to remove existing database: %v", err)
		}
	}

	// Create new database
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return fmt.Errorf("failed to create database: %v", err)
	}
	defer db.Close()

	// Read and execute migration
	migration, err := os.ReadFile(filepath.Join("db", "migrations", "0001_init.sql"))
	if err != nil {
		return fmt.Errorf("failed to read migration file: %v", err)
	}

	if _, err := db.Exec(string(migration)); err != nil {
		return fmt.Errorf("failed to execute migration: %v", err)
	}

	fmt.Println("Database initialized successfully")
	return nil
}

// Migrate runs database migrations
func Migrate() error {
	// TODO: Implement database migrations
	return nil
}

// Seed imports initial data into the database
func Seed() error {
	// TODO: Implement database seeding
	return nil
}

// Reset clears all data and reinitializes the database
func Reset() error {
	// TODO: Implement full reset
	return nil
}

// ResetHistory clears only study history
func ResetHistory() error {
	// TODO: Implement history reset
	return nil
} 