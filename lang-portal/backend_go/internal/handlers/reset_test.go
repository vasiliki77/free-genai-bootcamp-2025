package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/lang-portal/backend_go/internal/handlers"
	"github.com/lang-portal/backend_go/internal/models"
	"github.com/lang-portal/backend_go/internal/service"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestResetHistoryHandler(t *testing.T) {
	// 1. Setup
	db, err := models.NewDB(":memory:")
	if err != nil {
		t.Fatalf("Failed to initialize test database: %v", err)
	}
	defer db.Close()
	runMigrationsForTest(t, db)

	studyService := service.NewStudyService(db)
	resetHandler := handlers.NewResetHandler(studyService)

	router := gin.Default()
	router.POST("/api/reset_history", resetHandler.ResetHistoryHandler)

	// 2. Request
	req, _ := http.NewRequest("POST", "/api/reset_history", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	// 3. Assertions
	t.Logf("Response Body: %s", recorder.Body.String()) // Log response body
	assert.Equal(t, http.StatusOK, recorder.Code, "Expected status code to be 200 OK")
	// Add more assertions later to check response body content
}

func TestFullResetHandler(t *testing.T) {
	// 1. Setup
	db, err := models.NewDB(":memory:")
	if err != nil {
		t.Fatalf("Failed to initialize test database: %v", err)
	}
	defer db.Close()
	runMigrationsForTest(t, db)

	studyService := service.NewStudyService(db)
	resetHandler := handlers.NewResetHandler(studyService)

	router := gin.Default()
	router.POST("/api/full_reset", resetHandler.FullResetHandler)

	// 2. Request
	req, _ := http.NewRequest("POST", "/api/full_reset", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	// 3. Assertions
	t.Logf("Response Body: %s", recorder.Body.String()) // Log response body
	assert.Equal(t, http.StatusOK, recorder.Code, "Expected status code to be 200 OK")
	// Add more assertions later to check response body content
}
