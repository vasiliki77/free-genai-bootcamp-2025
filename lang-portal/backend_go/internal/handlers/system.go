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
	err := h.service.ResetHistory()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Study history has been reset",
	})
}

func (h *SystemHandler) FullReset(c *gin.Context) {
	err := h.service.FullReset()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Database has been reset to initial state",
	})
} 