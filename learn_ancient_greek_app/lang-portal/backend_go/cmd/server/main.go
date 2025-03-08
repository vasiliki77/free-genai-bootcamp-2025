package main

import (
	"log"
	"os"

	"backend_go/internal/handlers"
	"backend_go/internal/models"
	"backend_go/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors" // Make sure to add this import
)

func main() {
	// Get database path from environment
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "words.db"
	}

	// Initialize services
	wordService := service.NewWordService()
	groupService := service.NewGroupService()
	systemService := service.NewSystemService()
	dashboardService := service.NewDashboardService()
	studyService := service.NewStudyService()
	translationService := service.NewTranslationService() // Add this line

	// Initialize handlers
	wordHandler := handlers.NewWordHandler(wordService)
	groupHandler := handlers.NewGroupHandler(groupService)
	systemHandler := handlers.NewSystemHandler(systemService)
	dashboardHandler := handlers.NewDashboardHandler(dashboardService)
	studyHandler := handlers.NewStudyHandler(studyService)
	translationHandler := handlers.NewTranslationHandler(translationService) // Add this line

	// Setup router
	router := gin.Default()
	
	// Add CORS middleware
	router.Use(cors.Default())
	api := router.Group("/api")

	// Word routes
	api.GET("/words", wordHandler.GetWords)
	api.GET("/words/:id", wordHandler.GetWord)

	// Group routes
	api.GET("/groups", groupHandler.GetGroups)
	api.GET("/groups/:id", groupHandler.GetGroup)
	api.GET("/groups/:id/words", groupHandler.GetGroupWords)
	api.GET("/groups/:id/study_sessions", groupHandler.GetGroupStudySessions)

	// Add translation route
	api.POST("/translate", translationHandler.TranslateText)

	// Study routes
	api.GET("/study_activities", studyHandler.GetStudyActivities)
	api.GET("/study_activities/:id", studyHandler.GetStudyActivity)
	api.GET("/study_activities/:id/study_sessions", studyHandler.GetActivityStudySessions)
	api.GET("/study_sessions", studyHandler.GetStudySessions)
	api.GET("/study_sessions/:id", studyHandler.GetStudySession)
	api.GET("/study_sessions/:id/words", studyHandler.GetStudySessionWords)

	// Dashboard routes
	api.GET("/dashboard/last_study_session", dashboardHandler.GetLastStudySession)
	api.GET("/dashboard/study_progress", dashboardHandler.GetStudyProgress)
	api.GET("/dashboard/quick-stats", dashboardHandler.GetQuickStats)

	// System routes
	api.POST("/reset_history", systemHandler.ResetHistory)
	api.POST("/reload_test_data", systemHandler.ReloadTestData)
	api.POST("/full_reset", systemHandler.FullReset)

	// Start server
	if err := models.InitDB(dbPath); err != nil {
		log.Fatal(err)
	}

	// Make sure you have the cors package installed:
	// go get github.com/gin-contrib/cors
	
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
