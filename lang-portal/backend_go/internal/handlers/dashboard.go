package handlers

import (
	"net/http"

	"backend_go/internal/service"

	"github.com/gin-gonic/gin"
)

type DashboardHandler struct {
	service *service.DashboardService
}

func NewDashboardHandler(s *service.DashboardService) *DashboardHandler {
	return &DashboardHandler{service: s}
}

func (h *DashboardHandler) GetLastStudySession(c *gin.Context) {
	session, err := h.service.GetLastStudySession()
	if err != nil {
		if err.Error() == "no study sessions found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "No study sessions found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, session)
}

func (h *DashboardHandler) GetStudyProgress(c *gin.Context) {
	progress, err := h.service.GetStudyProgress()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, progress)
}

func (h *DashboardHandler) GetQuickStats(c *gin.Context) {
	stats, err := h.service.GetQuickStats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, stats)
} 