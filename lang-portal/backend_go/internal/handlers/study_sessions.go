package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lang-portal/backend_go/internal/service"
)

// StudySessionHandler holds the service for study session related operations.
type StudySessionHandler struct {
	studySessionService *service.StudySessionService
}

// NewStudySessionHandler creates a new StudySessionHandler.
func NewStudySessionHandler(studySessionService *service.StudySessionService) *StudySessionHandler {
	return &StudySessionHandler{studySessionService: studySessionService}
}

// GetStudySessionHandler handles the GET /api/study_sessions/:id endpoint.
func (h *StudySessionHandler) GetStudySessionHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid study session ID"})
		return
	}

	session, err := h.studySessionService.GetStudySession(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get study session"})
		return
	}
	c.JSON(http.StatusOK, session)
}

// ListStudySessionsHandler handles the GET /api/study_sessions endpoint.
func (h *StudySessionHandler) ListStudySessionsHandler(c *gin.Context) {
	// TODO: Implement pagination parameters from query
	page := 1
	pageSize := 100 // Default page size
	sessions, err := h.studySessionService.ListStudySessions(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list study sessions"})
		return
	}
	c.JSON(http.StatusOK, sessions)
}

// GetWordsInStudySessionHandler handles the GET /api/study_sessions/:id/words endpoint.
func (h *StudySessionHandler) GetWordsInStudySessionHandler(c *gin.Context) {
	idStr := c.Param("id")
	sessionID, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid study session ID"})
		return
	}
	// TODO: Implement pagination parameters from query
	page := 1
	pageSize := 100 // Default page size
	words, err := h.studySessionService.GetWordsInStudySession(sessionID, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get words in study session"})
		return
	}
	c.JSON(http.StatusOK, words)
}

// ReviewWordInStudySessionHandler handles the POST /api/study_sessions/:id/words/:word_id/review endpoint.
func (h *StudySessionHandler) ReviewWordInStudySessionHandler(c *gin.Context) {
	sessionIDStr := c.Param("id")
	sessionID, err := strconv.Atoi(sessionIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid study session ID"})
		return
	}
	wordIDStr := c.Param("word_id")
	wordID, err := strconv.Atoi(wordIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid word ID"})
		return
	}

	var requestBody struct {
		Correct bool `json:"correct" binding:"required"`
	}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	reviewItem, err := h.studySessionService.ReviewWordInStudySession(sessionID, wordID, requestBody.Correct)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to review word in study session"})
		return
	}
	c.JSON(http.StatusOK, reviewItem)
} 