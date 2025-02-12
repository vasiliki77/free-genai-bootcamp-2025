package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lang-portal/backend_go/internal/service"
)

// WordHandler holds the service for word related operations.
type WordHandler struct {
	wordService *service.WordService
}

// NewWordHandler creates a new WordHandler.
func NewWordHandler(wordService *service.WordService) *WordHandler {
	return &WordHandler{wordService: wordService}
}

// GetWordsHandler handles the GET /api/words endpoint.
func (h *WordHandler) GetWordsHandler(c *gin.Context) {
	// TODO: Implement pagination parameters from query
	page := 1
	pageSize := 100 // Default page size
	words, err := h.wordService.ListWords(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list words"})
		return
	}
	c.JSON(http.StatusOK, words)
}

// GetWordHandler handles the GET /api/words/:id endpoint.
func (h *WordHandler) GetWordHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid word ID"})
		return
	}

	word, err := h.wordService.GetWord(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get word"})
		return
	}
	c.JSON(http.StatusOK, word)
} 