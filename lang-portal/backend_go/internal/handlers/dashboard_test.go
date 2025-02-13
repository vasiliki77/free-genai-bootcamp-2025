package handlers_test

import (
	"encoding/json"
	"log"
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

func TestGetQuickStatsHandler(t *testing.T) {
	// 1. Setup
	db := handlers.SetupTest(t)
	defer handlers.TeardownTest(db)

	dashboardService := service.NewDashboardService(db)
	dashboardHandler := handlers.NewDashboardHandler(dashboardService)

	router := gin.Default()
	router.GET("/api/dashboard/quick-stats", dashboardHandler.GetQuickStatsHandler)

	// 2. Request
	req, _ := http.NewRequest("GET", "/api/dashboard/quick-stats", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	// 3. Assertions
	t.Logf("Response Body: %s", recorder.Body.String()) // Log response body
	assert.Equal(t, http.StatusOK, recorder.Code, "Expected status code to be 200 OK")
	// Add more assertions later to check response body content
}

func TestGetLastStudySessionHandler(t *testing.T) {
	// 1. Setup
	db := handlers.SetupTest(t)
	defer handlers.TeardownTest(db)
	var err error

	// Insert a group for the study session to reference
	insertGroupQuery := `
		INSERT INTO groups (name) VALUES (?)
	`
	_, err = db.Exec(insertGroupQuery, "Test Group")
	if err != nil {
		t.Fatalf("Failed to insert test group: %v", err)
	}
	log.Println("Successfully inserted test group")

	// Insert a study session for testing
	insertQuery := `
		INSERT INTO study_sessions (group_id, study_activity) VALUES (?, ?)
	`
	log.Printf("Executing insertion query: %s with args [%d, %d]\n", insertQuery, 1, 1) // Log insertion query
	_, err = db.Exec(insertQuery, 1, 1)
	if err != nil {
		t.Fatalf("Failed to insert study session: %v", err)
	}
	log.Println("Successfully inserted study session for test") // Add success log

	dashboardService := service.NewDashboardService(db)
	dashboardHandler := handlers.NewDashboardHandler(dashboardService)

	router := gin.Default()
	router.GET("/api/dashboard/last_study_session", dashboardHandler.GetLastStudySessionHandler)

	// 2. Request
	req, _ := http.NewRequest("GET", "/api/dashboard/last_study_session", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	// 3. Assertions
	t.Logf("Response Body: %s", recorder.Body.String())    // Log response body
	var response models.LastStudySession                   // Ensure correct type reference
	err = json.Unmarshal(recorder.Body.Bytes(), &response) // Use = for assignment
	if err != nil {
		t.Fatalf("Failed to parse JSON response: %v", err)
	}

	assert.NotNil(t, response, "Expected a study session in response")
	assert.Equal(t, 1, response.GroupID, "Expected GroupID to be 1") // Adjust assertions based on inserted data
	assert.Equal(t, 1, response.StudyActivityID, "Expected StudyActivityID to be 1")
	// Add more assertions to check other fields in the response
	assert.Equal(t, http.StatusOK, recorder.Code, "Expected status code to be 200 OK")
	// Add more assertions later to check response body content
}

func TestGetStudyProgressHandler(t *testing.T) {
	// 1. Setup
	db := handlers.SetupTest(t)
	defer handlers.TeardownTest(db)

	dashboardService := service.NewDashboardService(db)
	dashboardHandler := handlers.NewDashboardHandler(dashboardService)

	router := gin.Default()
	router.GET("/api/dashboard/study_progress", dashboardHandler.GetStudyProgressHandler)

	// 2. Request
	req, _ := http.NewRequest("GET", "/api/dashboard/study_progress", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	// 3. Assertions
	t.Logf("Response Body: %s", recorder.Body.String()) // Log response body
	assert.Equal(t, http.StatusOK, recorder.Code, "Expected status code to be 200 OK")
	// Add more assertions later to check response body content
}
