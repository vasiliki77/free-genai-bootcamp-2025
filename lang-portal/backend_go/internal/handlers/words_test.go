package handlers_test

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/lang-portal/backend_go/internal/handlers"
	"github.com/lang-portal/backend_go/internal/models"
	"github.com/lang-portal/backend_go/internal/service"
	_ "github.com/mattn/go-sqlite3"      // Import sqlite3 for testing
	"github.com/stretchr/testify/assert" // Using testify for assertions (optional, but helpful)
)

func TestGetWordsHandler(t *testing.T) {
	// 1. Setup: Initialize test database, services, handlers, and router
	db, err := models.NewDB(":memory:") // In-memory SQLite for testing
	if err != nil {
		t.Fatalf("Failed to initialize test database: %v", err)
	}
	defer db.Close()
	// Initialize database (run migrations, seed if needed - you'll need to adapt your InitDB/Seed logic for tests)
	runMigrationsForTest(t, db) // Assuming you create a helper function for migrations in tests

	wordService := service.NewWordService(db)
	wordHandler := handlers.NewWordHandler(wordService)

	router := gin.Default()
	router.GET("/api/words", wordHandler.GetWordsHandler)

	// 2. Make Request
	req, _ := http.NewRequest("GET", "/api/words", nil) // Create GET request to /api/words
	recorder := httptest.NewRecorder()                  // Response recorder
	router.ServeHTTP(recorder, req)                     // Serve the request to the router

	// 3. Assertions
	// Log the response body to debug
	t.Logf("Response Body: %s", recorder.Body.String())
	assert.Equal(t, http.StatusOK, recorder.Code, "Expected status code to be 200 OK") // Assert status code

	var response models.WordsResponse // Assuming your handler returns models.WordsResponse
	err = json.Unmarshal(recorder.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to parse JSON response: %v", err)
	}

	// Add more assertions to check the response body content
	// For example, if you expect an empty list initially:
	assert.NotNil(t, response.Items, "Expected 'items' to be present in response")
	assert.Empty(t, response.Items, "Expected empty word list initially")
	assert.NotNil(t, response.Pagination, "Expected 'pagination' to be present in response")
	assert.Equal(t, 1, response.Pagination.CurrentPage, "Expected current page to be 1")
	// ... more pagination assertions as needed
}

// Helper function to run migrations for tests (adapt your migration logic)
func runMigrationsForTest(t *testing.T, db *models.DB) {
	migration, err := os.ReadFile(filepath.Join("..", "..", "db", "migrations", "0001_init.sql")) // Adjust path to migration file
	if err != nil {
		t.Fatalf("Failed to read migration file: %v", err)
	}
	log.Printf("Migration SQL: %s", string(migration)) // Log migration SQL

	_, err = db.Exec(string(migration))
	if err != nil {
		log.Printf("Migration execution error: %v", err) // Log migration error
		t.Fatalf("Failed to execute migration: %v", err)
	}
	log.Println("Migrations executed successfully") // Log success
}
