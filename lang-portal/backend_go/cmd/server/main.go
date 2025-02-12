package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/lang-portal/backend_go/internal/handlers"
	"github.com/lang-portal/backend_go/internal/models"
	"github.com/lang-portal/backend_go/internal/service"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	r := gin.Default()

	// Initialize database
	db, err := models.NewDB("words.db")
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Initialize service
	dashboardService := service.NewDashboardService(db)
	wordService := service.NewWordService(db)
	groupService := service.NewGroupService(db)
	studyActivityService := service.NewStudyActivityService(db)
	studySessionService := service.NewStudySessionService(db)
	studyService := service.NewStudyService(db)

	// Initialize handlers
	dashboardHandler := handlers.NewDashboardHandler(dashboardService)
	wordHandler := handlers.NewWordHandler(wordService)
	groupHandler := handlers.NewGroupHandler(groupService)
	studyActivityHandler := handlers.NewStudyActivityHandler(studyActivityService)
	studySessionHandler := handlers.NewStudySessionHandler(studySessionService)
	resetHandler := handlers.NewResetHandler(studyService)

	api := r.Group("/api")
	{
		// Dashboard routes
		api.GET("/dashboard/last_study_session", dashboardHandler.GetLastStudySessionHandler)
		api.GET("/dashboard/study_progress", dashboardHandler.GetStudyProgressHandler)
		api.GET("/dashboard/quick-stats", dashboardHandler.GetQuickStatsHandler)

		// Study activities routes
		api.GET("/study_activities", studyActivityHandler.ListStudyActivitiesHandler)
		api.GET("/study_activities/:id", studyActivityHandler.GetStudyActivityHandler)
		api.GET("/study_activities/:id/study_sessions", studyActivityHandler.GetStudySessionsForActivityHandler)
		api.POST("/study_activities", studyActivityHandler.CreateStudyActivitySessionHandler)

		// Words routes
		api.GET("/words", wordHandler.GetWordsHandler)
		api.GET("/words/:id", wordHandler.GetWordHandler)

		// Groups routes
		api.GET("/groups", groupHandler.ListGroupsHandler)
		api.GET("/groups/:id", groupHandler.GetGroupHandler)
		api.GET("/groups/:id/words", groupHandler.GetWordsInGroupHandler)
		api.GET("/groups/:id/study_sessions", groupHandler.GetStudySessionsForGroupHandler)

		// Study sessions routes
		api.GET("/study_sessions", studySessionHandler.ListStudySessionsHandler)
		api.GET("/study_sessions/:id", studySessionHandler.GetStudySessionHandler)
		api.GET("/study_sessions/:id/words", studySessionHandler.GetWordsInStudySessionHandler)
		api.POST("/study_sessions/:id/words/:word_id/review", studySessionHandler.ReviewWordInStudySessionHandler)

		// Reset routes
		api.POST("/reset_history", resetHandler.ResetHistoryHandler)
		api.POST("/full_reset", resetHandler.FullResetHandler)
	}

	log.Fatal(r.Run(":8080"))
} 