package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"backend_go/internal/service"

	"github.com/gin-gonic/gin"
)

type WordHandler struct {
	service *service.WordService
}

func NewWordHandler(s *service.WordService) *WordHandler {
	return &WordHandler{service: s}
}

func (h *WordHandler) GetWords(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	fmt.Printf("GetWords called with page: %v\n", page) // Debug
	if err != nil || page < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page parameter"})
		return
	}

	perPage, err := strconv.Atoi(c.DefaultQuery("per_page", "100"))
	if err != nil || perPage < 1 || perPage > 100 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid per_page parameter"})
		return
	}

	words, err := h.service.GetWords(page, perPage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, words)
}

func (h *WordHandler) GetWord(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid word ID format"})
		return
	}

	word, err := h.service.GetWord(uint(id))
	if err != nil {
		if err.Error() == "word not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Word not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, word)
} 