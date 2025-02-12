package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lang-portal/backend_go/internal/service"
)

// ResetHandler holds the service for reset related operations.
type ResetHandler struct {
	studyService *service.StudyService
}

// NewResetHandler creates a new ResetHandler.
func NewResetHandler(studyService *service.StudyService) *ResetHandler {
	return &ResetHandler{studyService: studyService}
}

// ResetHistoryHandler handles the POST /api/reset_history endpoint.
func (h *ResetHandler) ResetHistoryHandler(c *gin.Context) {
	err := h.studyService.ResetHistory()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to reset history"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Study history has been reset"})
}

// FullResetHandler handles the POST /api/full_reset endpoint.
func (h *ResetHandler) FullResetHandler(c *gin.Context) {
	err := h.studyService.FullReset()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to perform full reset"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Database has been reset to initial state"})
} 