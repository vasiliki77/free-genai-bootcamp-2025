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

func TestListGroupsHandler(t *testing.T) {
	// 1. Setup
	db, err := models.NewDB(":memory:")
	if err != nil {
		t.Fatalf("Failed to initialize test database: %v", err)
	}
	defer db.Close()
	runMigrationsForTest(t, db)

	groupService := service.NewGroupService(db)
	groupHandler := handlers.NewGroupHandler(groupService)

	router := gin.Default()
	router.GET("/api/groups", groupHandler.ListGroupsHandler)

	// 2. Request
	req, _ := http.NewRequest("GET", "/api/groups", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	// 3. Assertions
	t.Logf("Response Body: %s", recorder.Body.String()) // Log response body
	assert.Equal(t, http.StatusOK, recorder.Code, "Expected status code to be 200 OK")
	// Add more assertions later to check response body content
}

func TestGetGroupHandler(t *testing.T) {
	// 1. Setup
	db, err := models.NewDB(":memory:")
	if err != nil {
		t.Fatalf("Failed to initialize test database: %v", err)
	}
	defer db.Close()
	runMigrationsForTest(t, db)

	groupService := service.NewGroupService(db)
	groupHandler := handlers.NewGroupHandler(groupService)

	router := gin.Default()
	router.GET("/api/groups/:id", groupHandler.GetGroupHandler)

	// 2. Request
	req, _ := http.NewRequest("GET", "/api/groups/1", nil) // Example ID = 1
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	// 3. Assertions
	t.Logf("Response Body: %s", recorder.Body.String()) // Log response body
	assert.Equal(t, http.StatusOK, recorder.Code, "Expected status code to be 200 OK")
	// Add more assertions later to check response body content
}

func TestGetWordsInGroupHandler(t *testing.T) {
	// 1. Setup
	db, err := models.NewDB(":memory:")
	if err != nil {
		t.Fatalf("Failed to initialize test database: %v", err)
	}
	defer db.Close()
	runMigrationsForTest(t, db)

	groupService := service.NewGroupService(db)
	groupHandler := handlers.NewGroupHandler(groupService)

	router := gin.Default()
	router.GET("/api/groups/:id/words", groupHandler.GetWordsInGroupHandler)

	// 2. Request
	req, _ := http.NewRequest("GET", "/api/groups/1/words", nil) // Example ID = 1
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	// 3. Assertions
	t.Logf("Response Body: %s", recorder.Body.String()) // Log response body
	assert.Equal(t, http.StatusOK, recorder.Code, "Expected status code to be 200 OK")
	// Add more assertions later to check response body content
}

func TestGetStudySessionsForGroupHandler(t *testing.T) {
	// 1. Setup
	db, err := models.NewDB(":memory:")
	if err != nil {
		t.Fatalf("Failed to initialize test database: %v", err)
	}
	defer db.Close()
	runMigrationsForTest(t, db)

	groupService := service.NewGroupService(db)
	groupHandler := handlers.NewGroupHandler(groupService)

	router := gin.Default()
	router.GET("/api/groups/:id/study_sessions", groupHandler.GetStudySessionsForGroupHandler)

	// 2. Request
	req, _ := http.NewRequest("GET", "/api/groups/1/study_sessions", nil) // Example ID = 1
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	// 3. Assertions
	t.Logf("Response Body: %s", recorder.Body.String()) // Log response body
	assert.Equal(t, http.StatusOK, recorder.Code, "Expected status code to be 200 OK")
	// Add more assertions later to check response body content
}
