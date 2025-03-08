package handlers

import (
	"net/http"
	"backend_go/internal/service"

	"github.com/gin-gonic/gin"
)

// TranslationHandler handles translation-related requests
type TranslationHandler struct {
	service *service.TranslationService
}

// NewTranslationHandler creates a new instance of TranslationHandler
func NewTranslationHandler(service *service.TranslationService) *TranslationHandler {
	return &TranslationHandler{
		service: service,
	}
}

// TranslateText handles requests to translate text
func (h *TranslationHandler) TranslateText(c *gin.Context) {
	var req struct {
		Text string `json:"text" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	result, err := h.service.Translate(req.Text)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Translation service error"})
		return
	}

	c.JSON(http.StatusOK, result)
}