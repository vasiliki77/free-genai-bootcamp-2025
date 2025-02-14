package handlers

import (
	"backend_go/internal/service"

	"github.com/gin-gonic/gin"
)

// Handler holds all the HTTP handlers and their dependencies
type Handler struct {
	dashboard *DashboardHandler
	words     *WordHandler
	study     *StudyHandler
}

// NewHandler creates a new Handler instance with all dependencies
func NewHandler(
	dashboardService *service.DashboardService,
	wordService *service.WordService,
	studyService *service.StudyService,
) *Handler {
	return &Handler{
		dashboard: NewDashboardHandler(dashboardService),
		words:     NewWordHandler(wordService),
		study:     NewStudyHandler(studyService),
	}
}

// Register registers all routes to the given router group
func (h *Handler) Register(router *gin.Engine) {
	api := router.Group("/api")
	{
		// Dashboard routes
		api.GET("/dashboard/last-session", h.dashboard.GetLastStudySession)
		api.GET("/dashboard/progress", h.dashboard.GetStudyProgress)
		api.GET("/dashboard/stats", h.dashboard.GetQuickStats)

		// Word routes
		api.GET("/words", h.words.GetWords)
		api.GET("/words/:id", h.words.GetWord)

		// Study routes
		api.GET("/study/activities", h.study.GetStudyActivities)
		api.GET("/study/activities/:id", h.study.GetStudyActivity)
		api.POST("/study/sessions", h.study.CreateStudySession)
		api.POST("/study/sessions/:id/words/:word_id/review", h.study.ReviewWord)
	}
}
