package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lang-portal/backend_go/internal/service"
)

// DashboardHandler holds the service for dashboard related operations.
type DashboardHandler struct {
	dashboardService *service.DashboardService
}

// NewDashboardHandler creates a new DashboardHandler.
func NewDashboardHandler(dashboardService *service.DashboardService) *DashboardHandler {
	return &DashboardHandler{dashboardService: dashboardService}
}

// GetLastStudySessionHandler handles the GET /api/dashboard/last_study_session endpoint.
func (h *DashboardHandler) GetLastStudySessionHandler(c *gin.Context) {
	session, err := h.dashboardService.GetLastStudySession()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get last study session"})
		return
	}
	c.JSON(http.StatusOK, session)
}

// GetStudyProgressHandler handles the GET /api/dashboard/study_progress endpoint.
func (h *DashboardHandler) GetStudyProgressHandler(c *gin.Context) {
	progress, err := h.dashboardService.GetStudyProgress()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get study progress"})
		return
	}
	c.JSON(http.StatusOK, progress)
}

// GetQuickStatsHandler handles the GET /api/dashboard/quick-stats endpoint.
func (h *DashboardHandler) GetQuickStatsHandler(c *gin.Context) {
	stats, err := h.dashboardService.GetQuickStats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get quick stats"})
		return
	}
	c.JSON(http.StatusOK, stats)
} 