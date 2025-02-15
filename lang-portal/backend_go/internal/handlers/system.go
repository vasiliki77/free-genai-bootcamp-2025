package handlers

import (
	"net/http"

	"backend_go/internal/service"

	"github.com/gin-gonic/gin"
)

type SystemHandler struct {
	service *service.SystemService
}

func NewSystemHandler(s *service.SystemService) *SystemHandler {
	return &SystemHandler{service: s}
}

func (h *SystemHandler) ResetHistory(c *gin.Context) {
	if err := h.service.ResetHistory(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Study history has been reset",
	})
}

func (h *SystemHandler) FullReset(c *gin.Context) {
	if err := h.service.FullReset(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "System has been fully reset",
	})
}

func (h *SystemHandler) ReloadTestData(c *gin.Context) {
	if err := h.service.ReloadTestData(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
} 