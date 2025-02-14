package handlers

import (
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
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "100"))

	words, err := h.service.GetWords(page, perPage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, words)
}

func (h *WordHandler) GetWord(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	word, err := h.service.GetWord(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, word)
} 