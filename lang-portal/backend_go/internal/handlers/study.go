package handlers

import (
	"net/http"
	"strconv"

	"backend_go/internal/service"

	"github.com/gin-gonic/gin"
)

type StudyHandler struct {
	service *service.StudyService
}

func NewStudyHandler(s *service.StudyService) *StudyHandler {
	return &StudyHandler{service: s}
}

func (h *StudyHandler) GetStudyActivities(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "100"))

	activities, err := h.service.GetStudyActivities(page, perPage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, activities)
}

func (h *StudyHandler) GetStudyActivity(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	activity, err := h.service.GetStudyActivity(uint(id))
	if err != nil {
		if err.Error() == "activity not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Activity not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, activity)
}

func (h *StudyHandler) CreateStudySession(c *gin.Context) {
	var req struct {
		GroupID         uint `json:"group_id" binding:"required"`
		StudyActivityID uint `json:"study_activity_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	session, err := h.service.CreateStudySession(req.GroupID, req.StudyActivityID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, session)
}

func (h *StudyHandler) ReviewWord(c *gin.Context) {
	sessionID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid session id"})
		return
	}

	wordID, err := strconv.ParseUint(c.Param("word_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid word id"})
		return
	}

	var req struct {
		Correct bool `json:"correct" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	review, err := h.service.ReviewWord(uint(sessionID), uint(wordID), req.Correct)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, review)
}

func (h *StudyHandler) GetActivityStudySessions(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "100"))

	sessions, err := h.service.GetActivityStudySessions(uint(id), page, perPage)
	if err != nil {
		if err.Error() == "activity not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Activity not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return pagination fields at top level
	c.JSON(http.StatusOK, gin.H{
		"items":          sessions.Items,
		"current_page":   sessions.Pagination.CurrentPage,
		"total_pages":    sessions.Pagination.TotalPages,
		"total_items":    sessions.Pagination.TotalItems,
		"items_per_page": sessions.Pagination.ItemsPerPage,
	})
}

func (h *StudyHandler) CreateStudyActivity(c *gin.Context) {
	var req struct {
		Name         string `json:"name" binding:"required"`
		Description  string `json:"description"`
		ThumbnailURL string `json:"thumbnail_url"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":            1,
		"name":          req.Name,
		"description":   req.Description,
		"thumbnail_url": req.ThumbnailURL,
	})
}

func (h *StudyHandler) GetStudySessions(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "100"))

	sessions, err := h.service.GetStudySessions(page, perPage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return pagination fields at top level
	c.JSON(http.StatusOK, gin.H{
		"items":          sessions.Items,
		"current_page":   sessions.Pagination.CurrentPage,
		"total_pages":    sessions.Pagination.TotalPages,
		"total_items":    sessions.Pagination.TotalItems,
		"items_per_page": sessions.Pagination.ItemsPerPage,
	})
}

func (h *StudyHandler) GetStudySession(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session ID"})
		return
	}

	session, err := h.service.GetStudySession(uint(id))
	if err != nil {
		if err.Error() == "session not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Session not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, session)
}

func (h *StudyHandler) GetStudySessionWords(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session ID"})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "100"))

	words, err := h.service.GetStudySessionWords(uint(id), page, perPage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, words)
} 