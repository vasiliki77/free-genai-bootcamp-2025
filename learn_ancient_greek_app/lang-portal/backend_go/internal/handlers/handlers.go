package handlers

import (
	"backend_go/internal/service"

	"github.com/gin-gonic/gin"
)

// Handlers holds all the HTTP handlers and their dependencies
type Handlers struct {
	dashboard *DashboardHandler
	words     *WordHandler
	study     *StudyHandler
	groups    *GroupHandler
	system    *SystemHandler
}

// NewHandlers creates a new Handlers instance with all dependencies
func NewHandlers(
	dashboardService *service.DashboardService,
	wordService *service.WordService,
	studyService *service.StudyService,
	groupService *service.GroupService,
	systemService *service.SystemService,
) *Handlers {
	return &Handlers{
		dashboard: NewDashboardHandler(dashboardService),
		words:     NewWordHandler(wordService),
		study:     NewStudyHandler(studyService),
		groups:    NewGroupHandler(groupService),
		system:    NewSystemHandler(systemService),
	}
}

// Register registers all routes to the given router group
func (h *Handlers) Register(router *gin.Engine) {
	api := router.Group("/api")
	{
		// Dashboard routes
		api.GET("/dashboard/last_study_session", h.dashboard.GetLastStudySession)
		api.GET("/dashboard/study_progress", h.dashboard.GetStudyProgress)
		api.GET("/dashboard/quick-stats", h.dashboard.GetQuickStats)

		// Word routes
		api.GET("/words", h.words.GetWords)
		api.GET("/words/:id", h.words.GetWord)

		// Group routes
		api.GET("/groups", h.groups.GetGroups)
		api.GET("/groups/:id", h.groups.GetGroup)
		api.GET("/groups/:id/words", h.groups.GetGroupWords)
		api.GET("/groups/:id/study_sessions", h.groups.GetGroupStudySessions)

		// Study routes
		api.GET("/study_activities", h.study.GetStudyActivities)
		api.GET("/study_activities/:id", h.study.GetStudyActivity)
		api.GET("/study_activities/:id/study_sessions", h.study.GetActivityStudySessions)
		api.POST("/study_activities", h.study.CreateStudyActivity)

		// Study sessions
		api.GET("/study_sessions", h.study.GetStudySessions)
		api.GET("/study_sessions/:id", h.study.GetStudySession)
		api.GET("/study_sessions/:id/words", h.study.GetStudySessionWords)
		api.POST("/study_sessions/:id/words/:word_id/review", h.study.ReviewWord)

		// System routes
		api.POST("/reset_history", h.system.ResetHistory)
		api.POST("/full_reset", h.system.FullReset)
	}
}
